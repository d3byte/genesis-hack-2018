import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Observable } from 'rxjs';
import {environment} from "../../environments/environment";

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const token: string = localStorage.getItem(environment.tokenKey);

    const headers: any = {};

    if (token) {
      headers.Authorization = `Bearer ${token}`;
    }

    const authReq = req.clone({ setHeaders: { ...headers } });

    return next.handle(authReq);
  }
}
