import { Injectable, Inject } from "@angular/core";
import { Observable } from "rxjs";
import { HttpClient } from "@angular/common/http";
import { URL_CONSTANTS } from "../../app.constants";

@Injectable({
  providedIn: "root",
})
export class NominationCriteriaService {
  constructor(
    @Inject("BASE_API_URL") private baseUrl: string,
    private httpClient: HttpClient
  ) {}

  public getNominationCriteriaSeasons(): Observable<any> {
    return this.httpClient.get(
      this.baseUrl + URL_CONSTANTS["NOMINATION_CRITERIAS"]
    );
  }

}
