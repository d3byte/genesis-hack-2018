import { coerceBooleanProperty } from '@angular/cdk/coercion';
import { ScrollDispatcher } from '@angular/cdk/scrolling';
import { Component, ElementRef, EventEmitter, HostBinding, HostListener, Input, NgZone, OnDestroy, OnInit, Output } from '@angular/core';

import { Observable, of, Subject } from 'rxjs';
import { delay, distinctUntilChanged, flatMap, map, scan, startWith, takeUntil, takeWhile } from 'rxjs/operators';

import { $animations } from './animations';

export type digAnimations =
  | 'landing'
  | 'pulse'
  | 'beat'
  | 'heartBeat'
  | 'fadeIn'
  | 'fadeInRight'
  | 'fadeInLeft'
  | 'fadeInUp'
  | 'fadeInDown'
  | 'zoomIn'
  | 'fadeOut'
  | 'fadeOutRight'
  | 'fadeOutLeft'
  | 'fadeOutDown'
  | 'fadeOutUp'
  | 'zoomOut';
export type digAnimateSpeed = 'slower' | 'slow' | 'normal' | 'fast' | 'faster';

export class digRect {
  constructor(readonly left: number, readonly top: number, readonly right: number, readonly bottom: number) {}

  public get width(): number {
    return this.right - this.left;
  }
  public get height(): number {
    return this.bottom - this.top;
  }
}

@Component({
  selector: '[digAnimate]',
  template: '<ng-content></ng-content>',
  animations: $animations,
})
export class AnimateComponentTrigger implements OnInit, OnDestroy {
  private readonly timings = { slower: '3s', slow: '2s', normal: '1s', fast: '500ms', faster: '300ms' };

  private replay$ = new Subject<boolean>();
  private dispose$ = new Subject<void>();

  constructor(private elm: ElementRef, private scroll: ScrollDispatcher, private zone: NgZone) {}

  private get idle() {
    return { value: 'idle' };
  }
  private get play() {
    return {
      value: this.animate,
      //delay: this.delay,
      params: {
        timing: this.timings[this.speed] || '1s',
      },
    };
  }

  /** Selects the animation to be played */
  @Input('digAnimate') public animate: digAnimations;

  /** Speeds up or slows down the animation */
  @Input() public speed: digAnimateSpeed = 'normal';

  @HostBinding('@animate')
  public trigger: string | {} = 'idle';

  /** Disables the animation */
  @Input('disabled') public set disableAnimation(value: boolean) {
    this.disabled = coerceBooleanProperty(value);
  }
  @HostBinding('@.disabled') public disabled = false;

  /** Emits at the end of the animation */
  @Output() public start = new EventEmitter<void>();
  @HostListener('@animate.start') public animationStart() {
    this.start.emit();
  }

  /** Emits at the end of the animation */
  @Output() public done = new EventEmitter<void>();
  @HostListener('@animate.done') public animationDone() {
    this.done.emit();
  }

  /** When true, keeps the animation idle until the next replay triggers */
  @Input('paused') public set pauseAnimation(value: boolean) {
    this.paused = coerceBooleanProperty(value);
  }
  public paused: boolean = false;

  /** When true, triggers the animation on element scrolling in the viewport */
  @Input('aos') public set enableAOS(value: boolean) {
    this.aos = coerceBooleanProperty(value);
  }
  public aos: boolean = false;

  /** When true, triggers the animation on element scrolling in the viewport */
  @Input('once') public set aosOnce(value: boolean) {
    this.once = coerceBooleanProperty(value);
  }
  public once: boolean = false;

  /** Specifies the amout of visibility triggering AOS */
  @Input() public threshold: number = 0.2;

  /** Replays the animation */
  @Input() public set replay(replay: any) {
    // Skips whenever the animation never triggered
    if (this.trigger === 'idle') return;

    // Re-triggers the animation again on request
    if (coerceBooleanProperty(replay)) {
      this.trigger = this.idle;
      this.replay$.next(true);
    }
  }

  ngOnInit() {
    this.subscribeToAnimateTrigger();
  }

  ngOnDestroy() {
    this.dispose();
  }

  private subscribeToAnimateTrigger(): void {
    // Triggers the animation based on the input flags
    this.animateTrigger(this.elm).subscribe(trigger => {
      // Triggers the animation to play or to idle
      this.trigger = trigger ? this.play : this.idle;
    });
  }

  private dispose() {
    this.dispose$.next();
    this.dispose$.complete();
  }

  // Triggers the animation
  private animateTrigger(elm: ElementRef<HTMLElement>): Observable<boolean> {
    return this.animateReplay().pipe(flatMap(trigger => (this.aos ? this.animateOnScroll(elm) : of(trigger))));
  }

  // Triggers the animation deferred
  private animateReplay(): Observable<boolean> {
    return this.replay$.pipe(takeUntil(this.dispose$), delay(0), startWith(!this.paused));
  }

  // Triggers the animation on scroll
  private animateOnScroll(elm: ElementRef<HTMLElement>): Observable<boolean> {
    // Returns an AOS observable
    return this.scroll.ancestorScrolled(elm, 100).pipe(
      // Makes sure to dispose on destroy
      takeUntil(this.dispose$),
      // Starts with initial element visibility
      startWith(!this.paused && this.visibility >= this.threshold),
      // Maps the scrolling to the element visibility value
      map(() => this.visibility),
      // Applies an hysteresys, so, to trigger the animation on based on the treshold while off on full invisibility
      scan<number, boolean>((result, visiblility) => visiblility >= this.threshold || (result ? visiblility > 0 : false)),
      // Distincts the resulting triggers
      distinctUntilChanged(),
      // Stop taking the first on trigger when aosOnce is set
      takeWhile(trigger => !trigger || !this.once, true),
      // Run NEXT within the angular zone to trigger change detection back on
      flatMap(
        trigger => new Observable<boolean>(observer => this.zone.run(() => observer.next(trigger))),
      ),
    );
  }

  // Computes the element visibility ratio
  private get visibility() {
    return this.intersectRatio(this.clientRect(this.elm), this.getScrollingArea(this.elm));
  }

  private intersectRatio(rect: digRect, cont: digRect): number {
    // Return 1.0 when the element is fully within its scroller container
    if (rect.left > cont.left && rect.top > cont.top && rect.right < cont.right && rect.bottom < cont.bottom) {
      return 1.0;
    }

    // Computes the intersection area otherwise
    const a = Math.round(rect.width * rect.height);
    const b = Math.max(0, Math.min(rect.right, cont.right) - Math.max(rect.left, cont.left));
    const c = Math.max(0, Math.min(rect.bottom, cont.bottom) - Math.max(rect.top, cont.top));

    // Returns the amount of visible area
    return Math.round(((b * c) / a) * 10) / 10;
  }

  // Returns the rectangular surface area of the element's scrolling container
  private getScrollingArea(elm: ElementRef<HTMLElement>): digRect {
    // Gets the cdkScolling container, if any
    const scroller = this.scroll.getAncestorScrollContainers(elm).pop();
    // Returns the element's most likely scrolling container area
    return scroller ? this.clientRect(scroller.getElementRef()) : this.windowRect();
  }

  // Element client bounding rect helper
  private clientRect(elm: ElementRef<HTMLElement>): digRect {
    const el = !!elm && elm.nativeElement;
    return !!el && el.getBoundingClientRect();
  }

  private windowRect(): digRect {
    return new digRect(0, 0, window.innerWidth, window.innerHeight);
  }
}
