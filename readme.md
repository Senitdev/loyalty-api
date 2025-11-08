# 🎯 LoyaltyCard – Plateforme de fidélité marchands & clients

## 📌 Description

**LoyaltyCard** est une application web **PWA (Progressive Web App)** permettant aux **marchands** de gérer des programmes de fidélité numériques.  
Les **clients** peuvent accumuler, consulter et échanger leurs points directement depuis leur compte.

L'application comprend :
- 🧑‍💼 Une interface **marchand** pour gérer les récompenses, les clients et les transactions.
- 👥 Un espace **client** pour suivre ses points et ses avantages.
- ⚙️ Un **backend Go (Gin)** rapide et sécurisé avec base de données SQL.
- 💎 Un **frontend Next.js** moderne, responsive et mobile-first, prêt à être installé comme PWA.

---

## 🏗️ Architecture du projet

loyaltycard/
│
├── backend/ # API Go (Gin)
│ ├── cmd /# main.go # Point d'entrée du serveur
│ ├── controller/ # Logique métier et endpoints
│ ├── internals/ models/ # Définition des entités (User, Client, Reward, etc.)
| |-- dto /# comprend les Dta transfert
│ ├── repository/ # Accès aux données via GORM
│ ├── services/ # Couche service pour la logique métier
│ ├── handlers/ # Routes Gin
│ ├── database/ # Variables d’environnement, connexion DB, CORS
│ └── go.mod / go.sum # Dépendances Go
│
└── frontend/ # App Next.js (React + TypeScript)
├── app/ # Structure Next.js 14 (app router)
│ ├── page.tsx # Page d'accueil
│ ├── merchant/ # Espace marchand (rewards, clients, dashboard)
| └── client/ # Comprend espace client
│ └── login/ # Page de connexion
├── public/manifest.json # Manifest PWA
├── public/icons/ # Icônes PWA
├── next.config.js # Config Next.js + next-pwa
└── package.json # Dépendances JS

---

## 🚀 Fonctionnalités principales

### 🎁 Marchands
- Créer / modifier / désactiver des **récompenses**
- Gérer les **clients**
- Octroyer ou retirer des **points de fidélité**
- Consulter le **solde de points** d’un client
- Tableau de bord avec les **10 derniers clients** et **transactions récentes**

### 👤 Clients
- Accumuler des points via QR code ou transaction
- Consulter son **solde actuel**
- Voir l’historique des **transactions**
- Échanger ses points contre des **récompenses**

### 💻 Application
- Interface **mobile-first** optimisée PWA
- Authentification **JWT**
- API REST sécurisée
- Mode **offline** grâce au service worker
- Installation sur **Android / iOS / Desktop**

---

## ⚙️ Installation et exécution

### 🧩 Prérequis
- **Go** ≥ 1.21
- **PostgreSQL**
- **Git**

---

### 🔧 1. Cloner le projet
```bash
git clone https://github.com/Senitdev/loyaltycard.git
cd loyaltycard

🧱 2. Lancer le backend (Go / Gin)
cd loyalty-api
go mod tidy
go run cmd/main.go

# Le serveur démarre sur http://localhost:9090
🌐 4. Configurer le CORS (Backend)
Dans main.go :
config := cors.Config{
    AllowOrigins: []string{"http://localhost:3000"},
    AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
}
r.Use(cors.New(config))

👨‍💻 Auteurs

Papa Toure  – Développeur Full Stack
GitHub
 https://www.linkedin.com/in/papa-toure-6b1287389/
 Portfolio: https://mon-portfolio-dusky.vercel.app/projects