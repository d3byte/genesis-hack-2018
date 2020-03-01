import { AnswerTypeEnum } from './enums/answer-type.enum';
import { ConfigInterface } from './interfaces/config.interface';

export const mockConfig: ConfigInterface = {
  id: 'wadowajdiajwdiajwd',
  title: 'Статистика по седушкам-бутылкам',
  publicToken: 'wwgdayhuijokwkadijuhyegt52216t7y28uiowakdjnhbbajnwk',
  creatorId: 'zalupawidnaudjakowdwbhajdkawdlwa',
  expirationDate: '2020-02-10T09:25:01.542Z',
  questions: [
    {
      question: 'Как сильно вы любите садиться на бутылку?',
      answerType: AnswerTypeEnum.RATE,
      options: {
        rateScales: ['Плохо', 'Отлично'],
      },
      answers: [
        {
          text: '😟',
          value: 1,
        },
        {
          text: '😃',
          value: 5,
        },
      ],
    },
    {
      question: 'А как какать?',
      answerType: AnswerTypeEnum.RATE,
      options: {
        rateScales: ['Плохо', 'Отлично'],
      },
      answers: [
        {
          text: '😟',
          value: 1,
        },
        {
          text: '😒',
          value: 2,
        },
        {
          text: '😐',
          value: 3,
        },
        {
          text: '🙂',
          value: 4,
        },
        {
          text: '😃',
          value: 5,
        },
      ],
    },
  ],
};
