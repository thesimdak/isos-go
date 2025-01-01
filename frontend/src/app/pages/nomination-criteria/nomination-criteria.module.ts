import { NgModule } from '@angular/core';
import { DropdownModule } from 'primeng/dropdown';
import { CardModule } from 'primeng/card';
import { ButtonModule } from 'primeng/button';
import { CommonModule } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';
import { NominationCriteriaComponent } from './nomination-criteria.component';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from '../../app-routing.module';

@NgModule({
  imports: [
    CardModule,
    DropdownModule,
    ButtonModule,
    CommonModule,
    BrowserModule,
    AppRoutingModule,
    FormsModule
  ],
  declarations: [
    NominationCriteriaComponent
  ],
  providers: [],
  exports: [NominationCriteriaComponent],
})
export class NominationCriteriaModule { }