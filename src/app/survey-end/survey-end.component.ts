import { Component, Input } from '@angular/core';

@Component({
  selector: 'survey-end',
  templateUrl: './survey-end.component.html',
  styleUrls: ['./survey-end.component.scss'],
})
export class SurveyEndComponent {
  @Input() public expired: boolean;
}
