import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';
import {DropdownModule} from 'primeng/dropdown';
import {TableModule} from 'primeng/table';
import { TopResultsComponent } from './top-results.component';
import { FormsModule } from '@angular/forms';

@NgModule({
    imports: [
      BrowserModule,
      CommonModule,
      DropdownModule,
      TableModule,
      FormsModule
              ],
  declarations: [
    TopResultsComponent
  ],
  providers: [],
  exports: [TopResultsComponent],
})
export class TopResultsModule { }