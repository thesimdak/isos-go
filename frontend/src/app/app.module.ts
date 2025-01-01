import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavigationModule } from './shared/components/navigation/navigation.module';
import { LandingModule } from './pages/landing/landing.module';
import { CompetitionListModule } from './pages/competition-list/competition-list.module';
import { environment } from '../environments/environment';
import { SeasonResolver } from './resolvers/season.resolver';
import { ResultsModule } from './pages/results/results.module';
import { CompetitionResolver } from './resolvers/competition.resolver';
import { CategoriesResolver } from './resolvers/categories.resolver';
import { TopResultsCategoriesResolver } from './resolvers/top-results-categories.resolver';
import { TopResultsModule } from './pages/top-results/top-results.module';
import { CompetitionsResolver } from './resolvers/competitions.resolver';
import { CompetitionManagementModule } from './pages/management/competition-management.module';
import { AllCategoriesResolver } from './resolvers/all-categories.resolver';
import { NominationCriteriaModule } from './pages/nomination-criteria/nomination-criteria.module';
import { HttpClientModule, provideHttpClient } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    NavigationModule,
    CompetitionListModule,
    CompetitionManagementModule,
    NominationCriteriaModule,
    LandingModule,
    ResultsModule,
    TopResultsModule
  ],
  providers: [
    { provide: "BASE_API_URL", useValue: environment.apiUrl },
    provideHttpClient(),
    SeasonResolver,
    CompetitionResolver,
    CategoriesResolver,
    AllCategoriesResolver,
    TopResultsCategoriesResolver,
    CompetitionsResolver
],
  bootstrap: [AppComponent],
  schemas: [CUSTOM_ELEMENTS_SCHEMA]
})
export class AppModule { }
