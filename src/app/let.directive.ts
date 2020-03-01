import { Directive, Input, OnChanges, TemplateRef, ViewContainerRef } from '@angular/core';

interface ILetContext<T> {
  digLet: T;
}

@Directive({
  selector: '[digLet]',
})
export class LetDirective<T> implements OnChanges {
  @Input() private digLet: T;

  private context: ILetContext<T> = { digLet: null };

  constructor(private vcr: ViewContainerRef, private templateRef: TemplateRef<ILetContext<T>>) {
    this.vcr.createEmbeddedView(this.templateRef, this.context);
  }

  ngOnChanges(): void {
    this.context.digLet = this.digLet;
  }
}
