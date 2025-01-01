import { Component } from '@angular/core';
import { SelectItem } from 'primeng/api';
import { ActivatedRoute } from '@angular/router';
import { Category } from '../../shared/model/category.interface';

@Component({
  selector: 'isos-nomination-criteria',
  standalone: false,
  templateUrl: './nomination-criteria.component.html',
  styleUrls: ['./nomination-criteria.component.css']
})
export class NominationCriteriaComponent {

  public selectedCategory: number | undefined;
  public categories: SelectItem[];

  constructor(private route: ActivatedRoute) {
    this.categories = this.route.snapshot.data['categories'].map((category: Category) => { return { label: category.label, value: category.id } });
  }

  public selectCategory($event: { value: number | undefined; }) {
    this.selectedCategory = $event.value;
   
  }

}
