import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output } from '@angular/core';

import { AnswerTypeEnum } from '../services/config/enums/answer-type.enum';
import { QuestionInterface } from '../services/config/interfaces/question.interface';
import { SelectedAnswerInterface } from '../services/state/interfaces/selected-answer.interface';

@Component({
  selector: 'survey-question',
  templateUrl: './question.component.html',
  styleUrls: ['./question.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class QuestionComponent {
  @Input() public question: QuestionInterface;
  @Input() public selectedAnswer: SelectedAnswerInterface;

  @Output() public answer: EventEmitter<SelectedAnswerInterface> = new EventEmitter();

  public answerTypes: typeof AnswerTypeEnum = AnswerTypeEnum;

  public propagateAnswer(answer: SelectedAnswerInterface): void {
    this.answer.emit(answer);
  }
}
