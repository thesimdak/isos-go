import { Injectable, Inject } from "@angular/core";
import { Observable } from "rxjs";
import { HttpClient } from "@angular/common/http";
import { URL_CONSTANTS } from "../../app.constants";

@Injectable({
  providedIn: "root",
})
export class CompetitionService {
  constructor(
    @Inject("BASE_API_URL") private baseUrl: string,
    private httpClient: HttpClient
  ) {}

  public getCompetitionsBySeason(season: string): Observable<any> {
    return this.httpClient.get(
      this.baseUrl + URL_CONSTANTS["COMPETITIONS"].replace(":season", season)
    );
  }

  public getCompetition(competitionId: number): Observable<any> {
    return this.httpClient.get(
      this.baseUrl +
        URL_CONSTANTS["COMPETITION"].replace(":competitionId", competitionId.toString())
    );
  }

  public getAllCategories(competitionId: number): Observable<any> {
    return this.httpClient.get(
      this.baseUrl +
        URL_CONSTANTS["CATEGORIES"].replace(":competitionId", competitionId.toString())
    );
  }

  public getAllCategoriesTopResult(): Observable<any> {
    return this.httpClient.get(this.baseUrl + URL_CONSTANTS["CATEGORIES_ALL"]);
  }

  public getAllCompetitions(): Observable<any> {
    return this.httpClient.get(this.baseUrl + URL_CONSTANTS["COMPETITIONS_ALL"]);
  }

  public delete(competitionId: number): Observable<any> {
    return this.httpClient.delete(this.baseUrl +  URL_CONSTANTS['DELETE_COMPETITION']
      .replace(':competitionId', competitionId.toString())
    );
  }
}
