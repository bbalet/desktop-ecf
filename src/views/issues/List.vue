<script setup>
import { storeToRefs } from 'pinia';

import { useIssuesStore } from '@/stores';

const issuesStore = useIssuesStore();
const { issues } = storeToRefs(issuesStore);

usersStore.getAll();
</script>

<template>
    <h1>Users</h1>
    <router-link to="/issues/add" class="btn btn-sm btn-success mb-2">Ajouter un incident</router-link>
    <table class="table table-striped">
        <thead>
            <tr>
                <th style="width: 30%">Titre</th>
                <th style="width: 30%">Description</th>
                <th style="width: 10%"></th>
            </tr>
        </thead>
        <tbody>
            <template v-if="issues.length">
                <tr v-for="issue in issues" :key="issue.id">
                    <td>{{ issue.title }}</td>
                    <td>{{ issue.description }}</td>
                    <td style="white-space: nowrap">
                        <router-link :to="`/issues/edit/${issue.id}`" class="btn btn-sm btn-primary mr-1">Modifier</router-link>
                        <button @click="issuesStore.delete(issue.id)" class="btn btn-sm btn-danger btn-delete-issue" :disabled="issue.isDeleting">
                            <span v-if="issue.isDeleting" class="spinner-border spinner-border-sm"></span>
                            <span v-else>Supprimer</span>
                        </button>
                    </td>
                </tr>
            </template>
            <tr v-if="issues.loading">
                <td colspan="4" class="text-center">
                    <span class="spinner-border spinner-border-lg align-center"></span>
                </td>
            </tr>
            <tr v-if="issues.error">
                <td colspan="4">
                    <div class="text-danger">Erreur en chargeant les incidents : {{issues.error}}</div>
                </td>
            </tr>            
        </tbody>
    </table>
</template>
