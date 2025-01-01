import { Resolve, ActivatedRouteSnapshot } from '@angular/router';
import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { ResultService } from '../shared/services/result.service';

@Injectable()
export class SeasonResolver implements Resolve<Observable<string[]>> {


  constructor(private resultService: ResultService) {
    this.resultService = resultService;
  }

  public resolve(route: ActivatedRouteSnapshot): Observable<string[]>  {
    return this.resultService.getSeasons();
  }
}
