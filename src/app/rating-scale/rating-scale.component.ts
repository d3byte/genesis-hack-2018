import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output } from '@angular/core';

import { BehaviorSubject } from 'rxjs';

import { AnswerInterface } from '../services/config/interfaces/answer.interface';
import { SelectedAnswerInterface } from '../services/state/interfaces/selected-answer.interface';

@Component({
  selector: 'survey-rating-scale',
  templateUrl: './rating-scale.component.html',
  styleUrls: ['./rating-scale.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RatingScaleComponent {
  @Input() public rateScales: [string?, string?] = [];
  @Input() public answers: AnswerInterface[] = [];
  @Input() public answeredItem: SelectedAnswerInterface;

  @Output() public answer: EventEmitter<SelectedAnswerInterface> = new EventEmitter();

  public nextActiveRate$: BehaviorSubject<number> = new BehaviorSubject(null);

  public handleHover(num: number): void {
    this.nextActiveRate$.next(num);
  }

  public handleSelect(answer: SelectedAnswerInterface): void {
    this.answer.emit(answer?.index === this.answeredItem?.index ? null : answer);
  }
}
