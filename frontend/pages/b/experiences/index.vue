<template>
    <div class="experience-form-container">
      <div class="experience-form-content">
        <h2 class="form-title">体験登録フォーム</h2>
        <form @submit.prevent="handleSubmit" class="experience-form">
          <div class="form-group">
            <label for="name">体験名</label>
            <input id="name" v-model="formData.name" type="text" required>
          </div>
  
          <div class="form-group">
            <label for="description">体験情報</label>
            <textarea id="description" v-model="formData.description" required></textarea>
          </div>
  
          <div class="form-group">
            <label for="details">体験の具体的な内容</label>
            <textarea id="details" v-model="formData.details" required></textarea>
          </div>
  
          <div class="form-group">
            <label for="capacity">対応人数</label>
            <input id="capacity" v-model="formData.capacity" type="number" required>
          </div>
  
          <div class="form-group">
            <label for="preparations">準備物</label>
            <textarea id="preparations" v-model="formData.preparations"></textarea>
          </div>
  
          <div class="form-group checkbox-group">
            <input id="furusatoNozei" v-model="formData.isFurusatoNozei" type="checkbox">
            <label for="furusatoNozei">ふるさと納税対応</label>
          </div>
  
          <div class="form-group">
            <label for="price">金額（円）</label>
            <input id="price" v-model="formData.price" type="number" required>
          </div>
  
          <div class="form-group">
            <label for="ageRange">対象年齢</label>
            <input id="ageRange" v-model="formData.ageRange" type="text" required placeholder="例: 10歳以上">
          </div>
  
          <div class="form-group">
            <label for="startDate">体験開始可能日</label>
            <input id="startDate" v-model="formData.startDate" type="date" required>
          </div>
  
          <div class="form-group">
            <label for="endDate">体験終了日</label>
            <input id="endDate" v-model="formData.endDate" type="date" required>
          </div>
          <!-- NOTE: ボタン押しても何も動かない -->
          <v-btn type="submit" class="submit-button" nuxt to="/b/home">登録する</v-btn>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  
  const formData = ref({
    name: '',
    description: '',
    details: '',
    capacity: '',
    preparations: '',
    isFurusatoNozei: false,
    price: '',
    ageRange: '',
    startDate: '',
    endDate: '',
  })

  </script>
  
  <style scoped>
  .experience-form-container {
    width: 100%;
    max-width: 600px;
    margin: 0 auto;
    padding: 20px;
  }
  
  .experience-form-content {
    background-color: #ffffff;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    padding: 20px;
  }
  
  .form-title {
    font-size: 24px;
    font-weight: bold;
    text-align: center;
    margin-bottom: 20px;
  }
  
  .experience-form {
    display: flex;
    flex-direction: column;
    gap: 15px;
  }
  
  .form-group {
    display: flex;
    flex-direction: column;
  }
  
  .form-group label {
    margin-bottom: 5px;
    font-weight: bold;
  }
  
  .form-group input,
  .form-group textarea {
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 16px;
  }
  
  .form-group textarea {
    min-height: 100px;
  }
  
  .checkbox-group {
    flex-direction: row;
    align-items: center;
    gap: 10px;
  }
  
  .checkbox-group input {
    width: auto;
  }
  
  .submit-button {
    background-color: #4a90e2;
    color: white;
    border: none;
    padding: 10px 15px;
    font-size: 18px;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s;
  }
  
  .submit-button:hover {
    background-color: #357abd;
  }
  </style>