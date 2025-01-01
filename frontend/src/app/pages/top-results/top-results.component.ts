import { Component, OnDestroy } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Observable, Subscription } from 'rxjs';
import { ResultItem } from '../../shared/model/result-item.interface';
import { SelectItem } from 'primeng/api';
import { ResultService } from '../../shared/services/result.service';
import { Category } from '../../shared/model/category.interface';

@Component({
  selector: 'isos-top-results',
  standalone: false,
  templateUrl: './top-results.component.html',
  styleUrls: ['./top-results.component.css']
})
export class TopResultsComponent implements OnDestroy {

  private subscription: Subscription;
  public selectedCategory: number;
  public categories: SelectItem[];
  public results: ResultItem[] = [];
  public loading: boolean = true;

  constructor(private route: ActivatedRoute, private resultService: ResultService) {
    this.categories = this.route.snapshot.data['categories'].map((category: Category) => { return {label: category.label, value: category.id} }) ;
    this.selectedCategory = this.categories[0].value;
    this.subscription = this.resultService.getTopResult(this.selectedCategory)
                .subscribe((results) => {
                  this.loading = false;
                  this.results = results;
                });
  }

  public selectCategory($event: { value: number; }) {
    this.selectedCategory = $event.value;
    this.loading = true;
    this.subscription = this.resultService.getTopResult(this.selectedCategory)
                .subscribe((results) => {
                  this.loading = false;
                  this.results = results;
                });
  }

  public isTime3(): boolean {
    if (this.results != null) {
      for (const result of this.results) {
        if (result.time3) {
          return true;
        }
      }
    }
    return false;
  }

  public isTime4(): boolean {
    if (this.results != null) {
      for (const result of this.results) {
        if (result.time4) {
          return true;
        }
      }
    }
    return false;
  }

  public getStringTime(time: Number): String {
    if (time === undefined || time == null) {
      return "";
    }
    if (time === 999) {
      return '-';
    } else {
      return time.toFixed(2).toString();
    }
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }

}
