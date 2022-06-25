
<style scoped>
  mark-cyan {
    background-color: #17EFE7;
    color: #000000;
  }
  mark-orange {
    background-color: #F7C00E;
    color: #000000;
  }
  ul {
    font-size: 70%;
  }
</style>

**Un module WebAssembly "WASI" est :**

- Sécurisé
- Polyglotte
- Rapide
- <mark-cyan>Léger</mark-cyan>

**Un module WebAssembly <mark-orange>ne peut pas</mark-orange> :**

- Accéder au système d’exploitation
- Accéder à la mémoire que le host ne lui a pas donnée
- Faire des requêtes sur le réseau
- Lire ou écrire dans des fichiers


---
