<template>
    <div class="flex flex-col items-center justify-center min-h-screen bg-gray-100">
        <h1 class="text-3xl font-bold mb-6">Регистрация</h1>
        <form @submit.prevent="handleSubmit" class="bg-white p-6 rounded shadow-md w-80">
            <div class="mb-4">
                <label for="login" class="block text-sm font-medium text-gray-700">Логин</label>
                <input v-model="form.login" type="text" id="login" required
                    class="mt-1 block w-full p-2 border border-gray-300 rounded" />
            </div>
            <div class="mb-4">
                <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
                <input v-model="form.email" type="email" id="email" required
                    class="mt-1 block w-full p-2 border border-gray-300 rounded" />
            </div>
            <div class="mb-4">
                <label for="password" class="block text-sm font-medium text-gray-700">Пароль</label>
                <input v-model="form.password" type="password" id="password" required
                    class="mt-1 block w-full p-2 border border-gray-300 rounded" />
            </div>
            <div class="mb-4">
                <label for="full_name" class="block text-sm font-medium text-gray-700">Полное имя</label>
                <input v-model="form.full_name" type="text" id="full_name" required
                    class="mt-1 block w-full p-2 border border-gray-300 rounded" />
            </div>
            <button type="submit"
                class="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600">Зарегистрироваться</button>
        </form>
    </div>
</template>

<script setup lang="ts">
import { client, register } from '../../client/services.gen'

client.setConfig({
    baseUrl: 'http://5.23.53.194:8000',
});

const form = ref({
    login: '',
    password: '',
    email: '',
    full_name: ''
});

const handleSubmit = async () => {
    try {
        const response = await register({
            client,
            body: form.value
        });
        console.log(response);
        // Здесь вы можете обработать успешную регистрацию, например, перенаправить пользователя
    } catch (error) {
        console.error('Ошибка регистрации:', error);
        // Здесь вы можете обработать ошибку, например, показать сообщение об ошибке
    }
};
</script>

<style scoped>
/* Добавьте дополнительные стили, если необходимо */
</style>