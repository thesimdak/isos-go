import { NgModule } from '@angular/core';
import { CompetitionListComponent } from './competition-list.component';
import {DropdownModule} from 'primeng/dropdown';
import {CardModule} from 'primeng/card';
import {ButtonModule} from 'primeng/button';
import { CommonModule } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from '../../app-routing.module';

@NgModule({
    imports: [
      CardModule,
      DropdownModule,
      ButtonModule,
      CommonModule,
      BrowserModule,
      AppRoutingModule
              ],
  declarations: [
    CompetitionListComponent
  ],
  providers: [],
  exports: [CompetitionListComponent],
})
export class CompetitionListModule { }