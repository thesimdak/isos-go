import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';
import { DatePipe } from '@angular/common';
import { SelectItem } from 'primeng/api';
import { Competition } from '../../shared/model/competition.interface';
import { CompetitionService } from '../../shared/services/competition.service';

@Component({
  selector: 'isos-competition-list',
  standalone: false,
  templateUrl: './competition-list.component.html',
  styleUrls: ['./competition-list.component.css']
})
export class CompetitionListComponent {

  public selectedSeason: string;
  public seasons: SelectItem[];
  public competitions$: Observable<Competition[]>;
  private datePipe = new DatePipe('en-US');

  constructor(private route: ActivatedRoute, private competitionService: CompetitionService) {
    this.seasons = this.route.snapshot.data['seasons'].map((season: string) => { return {label: 'Sezóna ' + season, value: season} }).reverse();
    this.selectedSeason = this.seasons[0].value;
    this.competitions$ = this.competitionService.getCompetitionsBySeason(this.selectedSeason);
  }

  public selectSeason($event: { value: string; }) {
    this.selectedSeason = $event.value;
    this.competitions$ = this.competitionService.getCompetitionsBySeason(this.selectedSeason);
  }

  public getDate(date: string): string | null {
    return this.datePipe.transform(new Date(date), 'dd.MM.yyyy');
  }

}
