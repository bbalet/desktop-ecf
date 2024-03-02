<script setup>
import { storeToRefs } from 'pinia';
import { useAuthStore, useTheatersStore } from '@/stores';

const authStore = useAuthStore();
const { user } = storeToRefs(authStore);

const theaters = useTheatersStore()

theaters.getAll()

</script>

<template>
    <div v-if="user">
        <h1>Bonjour {{user.firstName}}!</h1>
        <p>Bienvenue sur l'application de gestion des incidents de Cinephoria</p>
        <div v-if="theaters">
            <h2>Choisir un cinéma</h2>
            <ul>
                <li v-for="theater in theaters" :key="theater.id">
                    {{ theater.id }} - {{ theater.city }}
                </li>
            </ul>
        </div>
    </div>
</template>
