export enum PrioritéEnum {
    Critical,
    Importante,
    Sérieux,
    AnnyiraNemQvaFontos,
    Unwichtig
}

export interface Tâche{
    TâcheIdentificationNuméro : number;
    UtilisateurIdentificationNuméro : number;
    Nom : string;
    Description : string;
    Priorité : PrioritéEnum;
    Fini : boolean;
    Terme : Date;
    CrééEn : Date;
}