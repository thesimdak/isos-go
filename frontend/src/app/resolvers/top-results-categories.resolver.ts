import { ActivatedRouteSnapshot, Resolve } from '@angular/router';
import { CompetitionService } from '../shared/services/competition.service';
import { Observable } from 'rxjs';
import { Category } from '../shared/model/category.interface';
import { Injectable } from '@angular/core';

@Injectable()
export class TopResultsCategoriesResolver implements Resolve<Observable<Category[]>> {
  private competitionService: CompetitionService;

  constructor(competitionService: CompetitionService) {
    this.competitionService = competitionService;
  }

  public resolve(route: ActivatedRouteSnapshot): Observable<Category[]>  {
    return this.competitionService.getAllCategoriesTopResult();
  }
}
