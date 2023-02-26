<script lang="ts" setup>
import { onMounted, Ref, ref } from 'vue';
import TodosApi from '@/services/todo-api';
// Logo
import logo from '../assets/logo.svg';
import { Todo } from '@/api/api-client';

const inputTodo = ref('');
const isError = ref(false);
const todos: Ref<Todo[] | undefined> = ref();
const editTodoItem: Ref<Todo | undefined> = ref();

const submit = async () => {
  try {
    if (!editTodoItem.value) {
      await TodosApi.postTodo({ content: inputTodo.value });
    } else {
      editTodoItem.value.content = inputTodo.value;
      await TodosApi.updateTodo(editTodoItem.value);
      editTodoItem.value = undefined;
      inputTodo.value = '';
    }
    if (isError) isError.value = false;
    getTodos();
  } catch (error) {
    console.error(error);
    isError.value = true;
  }
};

const getTodos = async () => {
  try {
    const res = await TodosApi.getTodos();
    todos.value = res.data;
  } catch (error) {
    console.error(error);
  }
};

const deleteTodo = async (id: number) => {
  try {
    await TodosApi.deleteTodo(id);
    getTodos();
  } catch (error) {
    console.error(error);
  }
};

const editTodo = (item: Todo) => {
  editTodoItem.value = item;
  if (item.content) inputTodo.value = item.content;
};

onMounted(() => {
  getTodos();
});
</script>

<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img :src="logo" class="my-3" contain height="200" />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">TODOアプリ サンプル</h1>
        <v-responsive class="mx-auto" max-width="300">
          <v-text-field
            v-model="inputTodo"
            label="TODOを入力"
            variant="outlined"
          ></v-text-field>
        </v-responsive>
        <v-btn @click="submit" color="blue">{{
          editTodoItem ? '保存' : '作成'
        }}</v-btn>
        <p v-if="isError" class="text-red text-center font-weight-bold">
          作成に失敗しました
        </p>
        <v-responsive class="mx-auto" max-width="500">
          <v-list density="compact">
            <v-list-subheader>TODO List</v-list-subheader>

            <v-list-item
              v-for="(item, i) in todos"
              :key="i"
              active-color="primary"
            >
              <v-list-item-title class="text-left">{{
                item.content
              }}</v-list-item-title>
              <template v-slot:append>
                <v-icon :icon="'mdi-pencil'" @click="editTodo(item)"></v-icon>
                <v-icon
                  :icon="'mdi-close'"
                  @click="deleteTodo(item.id)"
                ></v-icon>
              </template>
            </v-list-item>
          </v-list>
        </v-responsive>
      </v-col>
    </v-row>
  </v-container>
</template>
