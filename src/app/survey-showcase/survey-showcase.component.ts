import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output } from '@angular/core';

import { ConfigInterface } from '../services/config/interfaces/config.interface';

@Component({
  selector: 'survey-showcase',
  templateUrl: './survey-showcase.component.html',
  styleUrls: ['./survey-showcase.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SurveyShowcaseComponent {
  @Input() public config: ConfigInterface;

  @Output() public start: EventEmitter<never> = new EventEmitter();

  public handleClick(): void {
    this.start.emit();
  }
}
