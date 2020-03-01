import { Component, OnInit } from '@angular/core';

@Component({
  selector: '[survey-btn]',
  template: '<ng-content></ng-content>',
  styleUrls: ['./btn.component.scss'],
})
export class BtnComponent implements OnInit {
  constructor() {}

  ngOnInit(): void {}
}
