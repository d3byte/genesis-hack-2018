import { QuestionInterface } from './question.interface';

export interface ConfigInterface {
  title: string;
  id: string;
  questions: QuestionInterface[];
  publicToken: string;
  creatorId: string;
  expirationDate: string;
}
