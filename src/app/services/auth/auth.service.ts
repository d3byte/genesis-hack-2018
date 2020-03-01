import {HttpClient, HttpResponse} from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {environment} from "../../../environments/environment";
import {map, tap} from "rxjs/operators";

export interface LoginResponseInterface {
  token: string;
}

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  constructor(private httpClient: HttpClient) {}

  login$(): Observable<string> {
    return this.httpClient
      .post(`${environment.apiUrl}/login`, {})
      .pipe(
        map((response: HttpResponse<LoginResponseInterface>) => response.body.token),
        tap((token: string) => {
          localStorage.setItem(environment.tokenKey, token);
        }),
      );
  }
}
