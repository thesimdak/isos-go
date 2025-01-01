import { Component, OnInit } from "@angular/core";
import { ActivatedRoute, Router } from "@angular/router";
import { ConfirmationService } from "primeng/api";
import { SelectItem } from "primeng/api";
import { Observable } from "rxjs";
import { Inject } from '@angular/core';
import { MessageService } from 'primeng/api';
import { DatePipe } from "@angular/common";
import { NominationCriteriaPopupComponent } from "./nomination-criteria-popup/nomination-criteria-popup.component";
import { CompetitionService } from "../../shared/services/competition.service";
import { NominationCriteriaService } from "../../shared/services/nomination-criteria.service";
import { DialogService } from "primeng/dynamicdialog";
import { AuthenticationService } from "../../shared/services/authentication.service";
import { Competition } from "../../shared/model/competition.interface";
import { URL_CONSTANTS } from "../../app.constants";

@Component({
  selector: "isos-competition-management",
  standalone: false,
  templateUrl: "./competition-management.component.html",
  styleUrls: ["./competition-management.component.css"],
})
export class CompetitionManagementComponent {
  public competitions: any[] = [];
  public url: string;
  public heslo: string = '';

  public selectedSeason: string = '';
  public seasons: SelectItem[] = [];
  public competitions$: Observable<Competition[]> = new Observable();;
  public nominationCriteriaYears$: Observable<number[]>  = new Observable();;
  private datePipe = new DatePipe('en-US');

  constructor( 
    @Inject("BASE_API_URL") private baseUrl: string,
    private route: ActivatedRoute,
    private router: Router,
    private competitionService: CompetitionService,
    private nominationCriteriaService: NominationCriteriaService,
    private dialogService: DialogService,
    public authenticationService: AuthenticationService,
    public confirmationService: ConfirmationService,
    private messageService: MessageService
  ) {
    this.url = this.baseUrl + URL_CONSTANTS["RESULT_UPLOAD"]
    this.init();
  }

  public init() {
    this.seasons = this.route.snapshot.data["seasons"]
      .map((season: string) => {
        return { label: "Sezóna " + season, value: season };
      })
      .reverse();
    this.selectedSeason = this.seasons[0].value;
    this.competitions$ = this.competitionService.getCompetitionsBySeason(
      this.selectedSeason
    );
    this.nominationCriteriaYears$ = this.nominationCriteriaService.getNominationCriteriaSeasons();
  }

  public selectSeason($event: { value: string; }) {
    this.selectedSeason = $event.value;
    this.competitions$ = this.competitionService.getCompetitionsBySeason(
      this.selectedSeason
    );
  }

  public onBasicUpload(): void {
    this.competitions$ = this.competitionService.getCompetitionsBySeason(
      this.selectedSeason
    );
  }

  public createNominationCriteria(): void {
    this.dialogService.open(NominationCriteriaPopupComponent, {
      header: 'Vytvořit nominační kriteria',
      width: '50%'
  });
  }

  public removeNominationCriteria($event: any, nominationCriteriaYear: any): void {

  }

  public onFileSend() {
    this.messageService.add({severity:'success', summary:'Uploaduje se...', detail:'Mějte trpělivost, upload závodu může asi minutku trvat.'});
 
  }

  public openResultList(competitionId: number): void {
    this.router.navigateByUrl("results/" + competitionId);
  }

  public authenticate(): void {
    this.authenticationService.authenticate(this.heslo);
  }

  public removeCompetition($event: { stopPropagation: () => void; }, competitionId: number): void {
    $event.stopPropagation();
    this.confirmationService.confirm({
      message: "Opravdu chcete vymazat závod?",
      accept: () => {
        this.competitionService
          .delete(competitionId)
          .subscribe(
            () =>
              (this.competitions$ = this.competitionService.getCompetitionsBySeason(
                this.selectedSeason
              ))
          );
          this.messageService.add({severity:'success', summary:'Závod bude odstraněn...', detail:'Mějte trpělivost, odstranění závodu může chvilku trvat.'});
      },
    });
  }

  public getDate(date: string): string | null {
    return this.datePipe.transform(new Date(date), 'dd.MM.yyyy');
  }
}
