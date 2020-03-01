import { Injectable } from '@angular/core';

import {EMPTY, Observable, of} from 'rxjs';

import { StateInterface } from './interfaces/state.interface';
import { mockState } from './state.mock';
import {ModelsHTTPError, ModelsStateInterface, StateService} from "../../swagger";
import {catchError, map} from "rxjs/operators";

@Injectable({
  providedIn: 'root',
})
export class SurveyStateService {
  private stateMock: StateInterface = mockState;

  constructor(private stateService: StateService) {}


  public getState(configId: string): Observable<StateInterface> {
    return new Observable(s => {
      this.stateService.stateIdGet(configId).pipe(
        map((r: any) => r.data),
        catchError((err: ModelsHTTPError) => {
            return of(mockState);
        })
      ).subscribe(state => {
        s.next(state);
        s.complete();
      }, state => {
        s.next(state);
        s.complete();
      })
    });
  }

  public setState(state: StateInterface): Observable<any> {
    return of(true);
  }

  public submitState(state: StateInterface, configId: string): Observable<any> {
    return this.stateService.stateIdPost(state, configId);
  }
}
