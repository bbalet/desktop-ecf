<script setup>
import { Form, Field } from 'vee-validate';
import * as Yup from 'yup';
import { useRoute } from 'vue-router';
import { storeToRefs } from 'pinia';

import { useIssuesStore, useAlertStore } from '@/stores';
import { router } from '@/router';

const issuesStore = useIssuesStore();
const alertStore = useAlertStore();
const route = useRoute();
const id = route.params.id;

let title = 'Ajouter un incident';
let user = null;
if (id) {
    // edit mode
    title = 'Modifier un incident';
    ({ user } = storeToRefs(usersStore));
    issuesStore.getById(id);
}

const schema = Yup.object().shape({
    title: Yup.string()
        .required('Le titre est obligatoire'),
    description: Yup.string()
        .required('La description est obligatoire')
});

async function onSubmit(values) {
    try {
        let message;
        if (user) {
            await issuesStore.update(user.value.id, values)
            message = 'Incident modifié';
        } else {
            await issuesStore.register(values);
            message = 'Incident ajouté';
        }
        await router.push('/issues');
        alertStore.success(message);
    } catch (error) {
        alertStore.error(error);
    }
}
</script>

<template>
    <h1>{{title}}</h1>
    <template v-if="!(issue?.loading || issue?.error)">
        <Form @submit="onSubmit" :validation-schema="schema" :initial-values="issue" v-slot="{ errors, isSubmitting }">
            <div class="form-row">
                <div class="form-group col">
                    <label>Titre</label>
                    <Field name="title" type="text" class="form-control" :class="{ 'is-invalid': errors.title }" />
                    <div class="invalid-feedback">{{ errors.title }}</div>
                </div>
                <div class="form-group col">
                    <label>Description</label>
                    <Field name="description" type="text" class="form-control" :class="{ 'is-invalid': errors.description }" />
                    <div class="invalid-feedback">{{ errors.description }}</div>
                </div>
            </div>
            <div class="form-group">
                <button class="btn btn-primary" :disabled="isSubmitting">
                    <span v-show="isSubmitting" class="spinner-border spinner-border-sm mr-1"></span>
                    Save
                </button>
                <router-link to="/issues" class="btn btn-link">Annuler</router-link>
            </div>
        </Form>
    </template>
    <template v-if="issue?.loading">
        <div class="text-center m-5">
            <span class="spinner-border spinner-border-lg align-center"></span>
        </div>
    </template>
    <template v-if="issue?.error">
        <div class="text-center m-5">
            <div class="text-danger">Erreur en chargeant l'incident: {{issue.error}}</div>
        </div>
    </template>
</template>
