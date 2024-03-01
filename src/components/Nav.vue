<script setup>
import { useAuthStore } from '@/stores';

const authStore = useAuthStore();
</script>

<template>
    <nav v-show="authStore.user" class="navbar navbar-expand navbar-dark bg-dark">
        <div class="navbar-nav">
            <router-link to="/" class="nav-item nav-link">Accueil</router-link>
            <router-link to="/issues" class="nav-item nav-link">Incidents</router-link>
            <button @click="authStore.logout()" class="btn btn-link nav-item nav-link">Se déconnecter</button>
        </div>
    </nav>
</template>
