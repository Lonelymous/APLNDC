import { Component, OnInit, Input } from '@angular/core';
import { TasksService } from '../tasks.service';
import { Tâche } from '../tâche';

@Component({
  selector: 'app-tache-afficher',
  templateUrl: './tache-afficher.component.html',
  styleUrls: ['./tache-afficher.component.css']
})
export class TacheAfficherComponent implements OnInit {
  taches: Tâche[] = [];
  
  constructor(private tâcheService: TasksService) { }

  ngOnInit(): void {
    this.afficherTâche();
  }

  afficherTâche(): void {
    this.tâcheService.obtenirDesTâches()
    .subscribe(tâche => this.taches = tâche)
  }
}
