import { ChangeDetectionStrategy, Component, Input, OnDestroy, OnInit } from '@angular/core';

import {BehaviorSubject, EMPTY, Observable, of, Subject} from 'rxjs';
import {catchError, debounceTime, filter, flatMap, takeUntil, tap} from 'rxjs/operators';

import { bothSideCollapse, verticalCollapse } from '../animation';
import { ConfigService } from '../services/config/config.service';
import { ConfigInterface } from '../services/config/interfaces/config.interface';
import { SelectedAnswerInterface } from '../services/state/interfaces/selected-answer.interface';
import { StateInterface } from '../services/state/interfaces/state.interface';
import { SurveyStateService } from '../services/state/state.service';

import * as moment from 'moment';
import * as qs from 'qs';
import {ModelsConfigInterface} from "../swagger";
import {TokenService} from "../services/token/token.service";

export const SELECTOR = 'micro-survey-widget';

@Component({
  selector: SELECTOR,
  templateUrl: './survey-widget.component.html',
  styleUrls: ['./survey-widget.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
  animations: [bothSideCollapse, verticalCollapse],
})
export class SurveyWidgetComponent implements OnInit, OnDestroy {
  private ngOnDestroy$: Subject<void> = new Subject();

  @Input() private configId: string;

  public config$: BehaviorSubject<ModelsConfigInterface> = new BehaviorSubject(null);
  public activeQuestionIndex$: BehaviorSubject<number> = new BehaviorSubject(null);

  public state$: BehaviorSubject<StateInterface> = new BehaviorSubject(null);

  public expired: boolean;

  constructor(private configService: ConfigService, private stateService: SurveyStateService, private tokenService: TokenService) {}

  ngOnInit() {
    this.getConfig();
  }

  ngOnDestroy() {
    this.ngOnDestroy$.next();
  }

  private getConfig(): void {
    this.tokenService
      .getToken()
      .pipe(
        flatMap(token => this.getConfigId()
          .pipe(
            flatMap((configId: string) => this.configService.getConfig(configId)),
            tap(config => this.config$.next(config)),
            filter(config => !this.checkExpire(config.expirationDate)),
            flatMap(config => this.stateService.getState(config.id)),
            tap((state: StateInterface) => {
              this.state$.next(state);
              this.activeQuestionIndex$.next(state?.activeQuestionIndex);
              // this.handleUpdateState();
            }),
            takeUntil(this.ngOnDestroy$),
          ))
      ).subscribe();
  }

  private checkExpire(expirationDate: string): boolean {
    this.expired = moment().isAfter(moment(expirationDate));
    return this.expired;
  }

  private saveState(): void {
    this.stateService.submitState(this.state$.value, this.configId).subscribe();
  }

  private getConfigId(): Observable<string> {
    if (!this.configId) {
      this.configId = qs.parse(window.location.search.slice(1)).configId;
    }

    return of(this.configId);
  }
  //
  // private handleUpdateState(): void {
  //   this.state$
  //     .pipe(
  //       debounceTime(300),
  //       flatMap((state: StateInterface) => this.stateService.setState(state)),
  //       takeUntil(this.ngOnDestroy$),
  //     )
  //     .subscribe();
  // }

  public start(): void {
    this.handleQuestionChange(0);
  }

  public handleQuestionChange(index: number): void {
    this.activeQuestionIndex$.next(index);
    this.state$.next({
      ...this.state$.value,
      activeQuestionIndex: index,
      completed: index === this.config$.value.questions.length,
    });

    if (index === this.config$.value.questions.length) {
      this.saveState();
    }
  }

  public handleAnswer(questionIndex: number, answer: SelectedAnswerInterface): void {
    const state = this.state$.value || {};
    const handledState = {
      ...state,
      answers: {
        ...state?.answers,
        [questionIndex]: {
          value: answer?.value,
          index: answer?.index,
        },
      },
    };

    if ([answer?.value, answer?.index].includes(undefined)) {
      delete handledState.answers[questionIndex];
    }

    this.state$.next(handledState);
  }
}
