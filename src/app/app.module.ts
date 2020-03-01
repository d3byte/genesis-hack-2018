import { Injector, NgModule } from '@angular/core';
import { createCustomElement } from '@angular/elements';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MatIconModule } from '@angular/material/icon';

import { AnimateModule } from './animate/animate.module';
import { BtnComponent } from './btn/btn.component';
import { PaginatorComponent } from './paginator/paginator.component';
import { QuestionComponent } from './question/question.component';
import { RatingScaleComponent } from './rating-scale/rating-scale.component';
import { SurveyEndComponent } from './survey-end/survey-end.component';
import { SurveyShowcaseComponent } from './survey-showcase/survey-showcase.component';
import { SELECTOR as SurveySelector, SurveyWidgetComponent } from './survey-widget/survey-widget.component';
import { LetDirective } from './let.directive';
import {environment} from "../environments/environment";
import {ApiModule, Configuration} from "./swagger";
import {HTTP_INTERCEPTORS, HttpClientModule} from "@angular/common/http";
import {AuthInterceptor} from "./interceptors/auth.interceptor";

export function GetApiConfiguration() {
  return new Configuration({
    apiKeys: {},
    basePath: environment.apiUrl,
  });
}

@NgModule({
  declarations: [
    SurveyWidgetComponent,
    RatingScaleComponent,
    RatingScaleComponent,
    BtnComponent,
    LetDirective,
    QuestionComponent,
    SurveyShowcaseComponent,
    PaginatorComponent,
    SurveyEndComponent,
  ],
  imports: [BrowserModule, HttpClientModule, BrowserAnimationsModule, MatIconModule, AnimateModule, ApiModule.forRoot(GetApiConfiguration)],
  providers: [{ provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true }],
  bootstrap: [SurveyWidgetComponent],
})
export class AppModule {
  constructor(private injector: Injector) {
    this.register(SurveyWidgetComponent, SurveySelector);
    this.setUpMaterialIcons();
  }

  private register(component: any, selector: string): void {
    const el = createCustomElement(component, { injector: this.injector });
    customElements.define(selector, el);
  }

  private setUpMaterialIcons(): void {
    // <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    const el = document.createElement('link');
    el.href = 'https://fonts.googleapis.com/icon?family=Material+Icons';
    el.rel = 'stylesheet';
    document.head.appendChild(el);
  }
}
