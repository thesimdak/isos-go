import { Component, OnInit } from '@angular/core';
import { Router, Event, NavigationStart, NavigationEnd, NavigationError, ActivatedRoute, RouterEvent, NavigationCancel } from '@angular/router';

@Component({
  selector: 'app-root',
  standalone: false,
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  public title = 'svetsplhu.cz - Výsledky';

  public loading: boolean = true;
  public showNav = false;

  constructor(private router: Router, 
    private activeRoute: ActivatedRoute) {
    
  }
  ngOnInit(): void {

    this.router.events.subscribe(evt => {
      // this is an injected Router instance
      if (evt instanceof NavigationEnd) {
        this.activeRoute.queryParams.subscribe(
          params => {
            if (params['resultView'] != undefined) {
              this.showNav = "true" !== params['resultView'];
            } else if (params  != undefined) {
              this.showNav = true;
            }
          }
        );
      }
    });
    
    this.router.events.subscribe((routerEvent: Event) => {
      this.checkRouterEvent(routerEvent);
    });
  }



  checkRouterEvent(routerEvent: Event): void {
    if (routerEvent instanceof NavigationStart) {
      this.loading = true;
    }

    if (routerEvent instanceof NavigationEnd ||
      routerEvent instanceof NavigationCancel ||
      routerEvent instanceof NavigationError) {
      this.loading = false;
    }
  }

}
