import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';
import {DropdownModule} from 'primeng/dropdown';
import { NominationCriteriaPopupComponent } from './nomination-criteria-popup.component';
import { ButtonModule } from 'primeng/button';
import { AppRoutingModule } from '../../../app-routing.module';

@NgModule({
    imports: [
      BrowserModule,
      CommonModule,
      DropdownModule,
      AppRoutingModule,
      ButtonModule
              ],
  declarations: [
    NominationCriteriaPopupComponent
  ],
  providers: [],
  exports: [NominationCriteriaPopupComponent],
})
export class NominationCriteriaPopupModule { }