import { AnswerTypeEnum } from '../enums/answer-type.enum';
import { InputTypeEnum } from '../enums/input-type.enum';

import { AnswerInterface } from './answer.interface';

export interface QuestionInterface {
  question: string;
  answerType: AnswerTypeEnum;
  answers?: AnswerInterface[];
  options?: {
    inputType?: InputTypeEnum;
    rateScales?: [string, string];
    optional?: boolean;
  };
}
