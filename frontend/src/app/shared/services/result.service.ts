import { Injectable, Inject } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { URL_CONSTANTS } from '../../app.constants';

@Injectable({
    providedIn: 'root'
})
export class ResultService {

    constructor(@Inject('BASE_API_URL') private baseUrl: string,
        private httpClient: HttpClient) { }


    public getSeasons(): Observable<any> {
        return this.httpClient.get(this.baseUrl + URL_CONSTANTS['SEASONS']);
    }

    public getResult(competitionId: number, categoryId: number): Observable<any> {
        return this.httpClient.get(this.baseUrl + URL_CONSTANTS['RESULT_LIST']
            .replace(':competitionId', competitionId.toString())
            .replace(':categoryId', categoryId.toString())
        );
    }

    public getTopResult(categoryId: number): Observable<any> {
        return this.httpClient.get(this.baseUrl + URL_CONSTANTS['TOP_RESULT_LIST']
            .replace(':categoryId', categoryId.toString())
        );
    }
}