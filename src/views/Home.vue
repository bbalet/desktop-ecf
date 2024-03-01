<script setup>
import { storeToRefs } from 'pinia';

import { useAuthStore } from '@/stores';

const authStore = useAuthStore();
const { user } = storeToRefs(authStore);
</script>

<template>
    <div v-if="user">
        <h1>Bonjour {{user.firstName}}!</h1>
        <p>Bienvenue sur l'application de gestion des incidents de Cinephoria</p>
        <p><router-link to="/issues">Voir les incidents</router-link></p>
    </div>
</template>
