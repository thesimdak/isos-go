import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { SeasonResolver } from './resolvers/season.resolver';
import { LandingComponent } from './pages/landing/landing.component';
import { CompetitionListComponent } from './pages/competition-list/competition-list.component';
import { ResultsComponent } from './pages/results/results.component';
import { CompetitionResolver } from './resolvers/competition.resolver';
import { CategoriesResolver } from './resolvers/categories.resolver';
import { TopResultsCategoriesResolver } from './resolvers/top-results-categories.resolver';
import { TopResultsComponent } from './pages/top-results/top-results.component';
import { CompetitionManagementModule } from './pages/management/competition-management.module';
import { CompetitionsResolver } from './resolvers/competitions.resolver';
import { CompetitionManagementComponent } from './pages/management/competition-management.component';
import { NominationCriteriaComponent } from './pages/nomination-criteria/nomination-criteria.component';
import { AllCategoriesResolver } from './resolvers/all-categories.resolver';

const routes: Routes = [
  {
    path: 'competition-list',
    component: CompetitionListComponent,
    resolve: { seasons: SeasonResolver }
  },
  {
    path: 'results/:resultId',
    component: ResultsComponent,
    resolve: { competition: CompetitionResolver, categories: CategoriesResolver }
  },
  {
    path: 'top-results',
    component: TopResultsComponent,
    resolve: { categories: TopResultsCategoriesResolver }
  },
  {
    path: 'management',
    component: CompetitionManagementComponent,
    resolve: { seasons: SeasonResolver }
  },
  {
    path: '',
    component: LandingComponent
  },
  {
    path: 'nomination-criteria/:season',
    component: NominationCriteriaComponent,
    resolve: { categories: AllCategoriesResolver }
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
