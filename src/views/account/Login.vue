<script setup>
import { Form, Field } from 'vee-validate';
import * as Yup from 'yup';

import { useAuthStore } from '@/stores';

const schema = Yup.object().shape({
    email: Yup.string().required("l'email est requis"),
    password: Yup.string().required('le mot de passe est requis')
});

async function onSubmit(values) {
    const authStore = useAuthStore();
    const { email, password } = values;
    await authStore.login(email, password);
}
</script>

<template>
    <div class="card m-3">
        <h4 class="card-header">Connexion</h4>
        <div class="card-body">
            <Form @submit="onSubmit" :validation-schema="schema" v-slot="{ errors, isSubmitting }">
                <div class="form-group">
                    <label>E-mail</label>
                    <Field name="email" type="text" class="form-control" :class="{ 'is-invalid': errors.username }" />
                    <div class="invalid-feedback">{{ errors.username }}</div>
                </div>
                <div class="form-group">
                    <label>Mot de passe</label>
                    <Field name="password" type="password" class="form-control" :class="{ 'is-invalid': errors.password }" />
                    <div class="invalid-feedback">{{ errors.password }}</div>
                </div>
                <div class="form-group">
                    <button class="btn btn-primary" :disabled="isSubmitting">
                        <span v-show="isSubmitting" class="spinner-border spinner-border-sm mr-1"></span>
                        Login
                    </button>
                </div>
            </Form>
        </div>
    </div>
</template>
