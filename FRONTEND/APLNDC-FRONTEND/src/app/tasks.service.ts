import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Tâche, PrioritéEnum } from './tâche';

@Injectable({
  providedIn: 'root'
})
export class TasksService {

  constructor() { }

  obtenirDesTâches(): Observable<Tâche[]> {
    let lisTâches: Tâche[] = [];

    var monEntêtes = new Headers();
    monEntêtes.append("Content-Type", "application/json");

    var demandeOptions = {
      method: 'GET',
      headers: monEntêtes,
    };

    /*valeurs de test*/
    let tâcheNouveau: Tâche = {
      TâcheIdentificationNuméro: 0,
      UtilisateurIdentificationNuméro: 0,
      Nom: "aaaaaaaaaaaaaaaaaaa",
      Description: "Lorem ipsum suck my balls and nuts :))))",
      Priorité: PrioritéEnum.Critical,
      Fini: true,
      Terme: new Date(),
      CrééEn: new Date()
    } 
    
    for(let i = 0; i < 10; i++){
      lisTâches.push(tâcheNouveau);
    }

    fetch("http://81.182.202.18:4000/tasks", demandeOptions)
      .then(response => {
        response.json().then(x => {
          lisTâches = [];

          let values: any[] = x;

          for (let i = 0; i < values.length; i++) {
            let obj = values[i];
            let tâcheNouveau: Tâche = {
              TâcheIdentificationNuméro: obj.tareaIdentificaciónNúmero,
              UtilisateurIdentificationNuméro: obj.usuarioIdentificaciónNúmero,
              Nom: obj.nombre,
              Description: obj.descripción,
              Priorité: obj.prioridad,
              Fini: obj.hecho,
              Terme: obj.plazo,
              CrééEn: obj.creadoEn
            }
            lisTâches.push(tâcheNouveau);
          }
          console.log(values)
        })
      })
      .then(result => console.log(result))
      .catch(error => {
      });

    return of(lisTâches);
  }
}
