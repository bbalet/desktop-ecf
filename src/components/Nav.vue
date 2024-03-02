<script setup>
import { useAuthStore } from '@/stores';

const authStore = useAuthStore();
</script>

<template>
    <nav v-show="authStore.user" class="navbar navbar-expand navbar-dark bg-dark">
        <div class="navbar-nav">
            <router-link to="/" class="navbar-brand"><img src="/logox48.png" />&nbsp;<strong>Cinéphoria</strong></router-link>
            <router-link to="/" class="nav-item nav-link">Accueil</router-link>
            <button @click="authStore.logout()" class="btn btn-link nav-item nav-link">Se déconnecter</button>
        </div>
    </nav>
</template>
