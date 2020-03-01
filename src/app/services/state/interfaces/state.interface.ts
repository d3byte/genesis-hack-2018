import { SelectedAnswerInterface } from './selected-answer.interface';

export interface StateInterface {
  answers?: {
    [index: number]: SelectedAnswerInterface;
  };
  activeQuestionIndex?: number;
  completed?: boolean;
}
