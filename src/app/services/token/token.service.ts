import { Injectable } from '@angular/core';

import {BehaviorSubject, Observable, of, ReplaySubject} from 'rxjs';

import { environment } from '../../../environments/environment';
import {UsersService} from "../../swagger";
import {map, tap} from "rxjs/operators";

@Injectable({
  providedIn: 'root',
})
export class TokenService {
  public token: string;

  constructor(private usersService: UsersService) { }

  public getToken(): Observable<string> {
    const token: string = localStorage.getItem(environment.tokenKey);

    let token$: Observable<string>;

    if (token === 'undefined') {
      // Get token from backend
      token$ = this.usersService
        .loginPost({} as any)
        .pipe(
          map((r: any) => r.data.token),
          tap(token => localStorage.setItem(environment.tokenKey, token))
        );
    } else {
      token$ = of(token);
    }

    return token$.pipe(
      tap(token => this.token = token),
    );
  }
}
