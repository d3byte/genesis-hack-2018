/**
 * MIPT Service API
 * This is MIPT microservice
 *
 * OpenAPI spec version: 0.1
 * Contact: support@digitory.dev
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *//* tslint:disable:no-unused-variable member-ordering */

import { Inject, Injectable, Optional }                      from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams,
         HttpResponse, HttpEvent }                           from '@angular/common/http';
import { CustomHttpUrlEncodingCodec }                        from '../encoder';

import { Observable }                                        from 'rxjs';

import { ModelsHTTPError } from '../model/modelsHTTPError';
import { ModelsHTTPSuccess } from '../model/modelsHTTPSuccess';
import { ModelsSpreadsheet } from '../model/modelsSpreadsheet';
import { ModelsSpreadsheetClear } from '../model/modelsSpreadsheetClear';

import { BASE_PATH, COLLECTION_FORMATS }                     from '../variables';
import { Configuration }                                     from '../configuration';


@Injectable()
export class SpreadsheetService {

    protected basePath = '//localhost:8081/api/v1';
    public defaultHeaders = new HttpHeaders();
    public configuration = new Configuration();

    constructor(protected httpClient: HttpClient, @Optional()@Inject(BASE_PATH) basePath: string, @Optional() configuration: Configuration) {
        if (basePath) {
            this.basePath = basePath;
        }
        if (configuration) {
            this.configuration = configuration;
            this.basePath = basePath || configuration.basePath || this.basePath;
        }
    }

    /**
     * @param consumes string[] mime-types
     * @return true: consumes contains 'multipart/form-data', false: otherwise
     */
    private canConsumeForm(consumes: string[]): boolean {
        const form = 'multipart/form-data';
        for (const consume of consumes) {
            if (form === consume) {
                return true;
            }
        }
        return false;
    }


    /**
     * Append Data
     * Append Data
     * @param body Fields are required
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public spreadsheetAppendPost(body: ModelsSpreadsheet, observe?: 'body', reportProgress?: boolean): Observable<ModelsHTTPSuccess>;
    public spreadsheetAppendPost(body: ModelsSpreadsheet, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<ModelsHTTPSuccess>>;
    public spreadsheetAppendPost(body: ModelsSpreadsheet, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<ModelsHTTPSuccess>>;
    public spreadsheetAppendPost(body: ModelsSpreadsheet, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (body === null || body === undefined) {
            throw new Error('Required parameter body was null or undefined when calling spreadsheetAppendPost.');
        }

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
            'application/json'
        ];
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.request<ModelsHTTPSuccess>('post',`${this.basePath}/spreadsheet/append`,
            {
                body: body,
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Clear Data
     * Clear Data
     * @param body Fields are required
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public spreadsheetClearPost(body: ModelsSpreadsheetClear, observe?: 'body', reportProgress?: boolean): Observable<ModelsHTTPSuccess>;
    public spreadsheetClearPost(body: ModelsSpreadsheetClear, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<ModelsHTTPSuccess>>;
    public spreadsheetClearPost(body: ModelsSpreadsheetClear, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<ModelsHTTPSuccess>>;
    public spreadsheetClearPost(body: ModelsSpreadsheetClear, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (body === null || body === undefined) {
            throw new Error('Required parameter body was null or undefined when calling spreadsheetClearPost.');
        }

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
            'application/json'
        ];
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.request<ModelsHTTPSuccess>('post',`${this.basePath}/spreadsheet/clear`,
            {
                body: body,
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Copy Google Table
     * Copy Google Table
     * @param body Fields are required
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public spreadsheetCopyPost(body: ModelsSpreadsheet, observe?: 'body', reportProgress?: boolean): Observable<ModelsHTTPSuccess>;
    public spreadsheetCopyPost(body: ModelsSpreadsheet, observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<ModelsHTTPSuccess>>;
    public spreadsheetCopyPost(body: ModelsSpreadsheet, observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<ModelsHTTPSuccess>>;
    public spreadsheetCopyPost(body: ModelsSpreadsheet, observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        if (body === null || body === undefined) {
            throw new Error('Required parameter body was null or undefined when calling spreadsheetCopyPost.');
        }

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
            'application/json'
        ];
        const httpContentTypeSelected: string | undefined = this.configuration.selectHeaderContentType(consumes);
        if (httpContentTypeSelected != undefined) {
            headers = headers.set('Content-Type', httpContentTypeSelected);
        }

        return this.httpClient.request<ModelsHTTPSuccess>('post',`${this.basePath}/spreadsheet/copy`,
            {
                body: body,
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

    /**
     * Create Google Table
     * Create Google Table
     * @param observe set whether or not to return the data Observable as the body, response or events. defaults to returning the body.
     * @param reportProgress flag to report request and response progress.
     */
    public spreadsheetPost(observe?: 'body', reportProgress?: boolean): Observable<ModelsHTTPSuccess>;
    public spreadsheetPost(observe?: 'response', reportProgress?: boolean): Observable<HttpResponse<ModelsHTTPSuccess>>;
    public spreadsheetPost(observe?: 'events', reportProgress?: boolean): Observable<HttpEvent<ModelsHTTPSuccess>>;
    public spreadsheetPost(observe: any = 'body', reportProgress: boolean = false ): Observable<any> {

        let headers = this.defaultHeaders;

        // to determine the Accept header
        let httpHeaderAccepts: string[] = [
            'application/json'
        ];
        const httpHeaderAcceptSelected: string | undefined = this.configuration.selectHeaderAccept(httpHeaderAccepts);
        if (httpHeaderAcceptSelected != undefined) {
            headers = headers.set('Accept', httpHeaderAcceptSelected);
        }

        // to determine the Content-Type header
        const consumes: string[] = [
        ];

        return this.httpClient.request<ModelsHTTPSuccess>('post',`${this.basePath}/spreadsheet`,
            {
                withCredentials: this.configuration.withCredentials,
                headers: headers,
                observe: observe,
                reportProgress: reportProgress
            }
        );
    }

}