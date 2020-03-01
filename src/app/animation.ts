import { animate, AnimationTriggerMetadata, state, style, transition, trigger } from '@angular/animations';

export const bothSideCollapse: AnimationTriggerMetadata = trigger('bothSideCollapse', [
  state(
    '*',
    style({
      overflow: 'hidden',
      height: '*',
      width: '*',
    }),
  ),

  state(
    'void',
    style({
      height: '0',
      width: '0',
      overflow: 'hidden',
    }),
  ),

  transition('* => void', animate('300ms ease-in-out')),

  transition('void => *', animate('300ms ease-in-out')),
]);

export const verticalCollapse: AnimationTriggerMetadata = trigger('verticalCollapse', [
  state(
    '*',
    style({
      overflow: 'hidden',
      height: '*',
    }),
  ),

  state(
    'void',
    style({
      height: '0',
      overflow: 'hidden',
    }),
  ),

  transition('* => void', animate('300ms ease-in-out')),

  transition('void => *', animate('300ms ease-in-out')),
]);
