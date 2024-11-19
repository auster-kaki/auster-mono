<template>
  <v-container>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="6">
        <h1 class="text-center mb-6">設定一覧</h1>

        <v-select
          v-model="selectedUser"
          :items="users"
          item-text="name"
          item-value="id"
          label="ユーザー切り替え"
          outlined
          @change="loadUserSettings"
        ></v-select>

        <v-form ref="form" v-model="valid" @submit.prevent="updateSettings">
          <v-text-field
            v-model="settings.name"
            label="ユーザー名"
            required
          ></v-text-field>

          <v-radio-group v-model="settings.gender" row>
            <v-radio label="男性" value="male"></v-radio>
            <v-radio label="女性" value="female"></v-radio>
            <v-radio label="その他" value="other"></v-radio>
          </v-radio-group>

          <v-text-field
            v-model.number="settings.age"
            label="年齢"
            type="number"
            required
          ></v-text-field>

          <v-combobox
            v-model="settings.hobbies"
            :items="hobbyOptions"
            item-text="name"
            label="趣味"
            multiple
            chips
            small-chips
            deletable-chips
          ></v-combobox>

          <v-file-input
            v-model="settings.photo"
            accept="image/*"
            label="顔写真"
            prepend-icon="mdi-camera"
          ></v-file-input>

          <v-row>
            <v-col cols="6">
              <v-btn
                color="primary"
                class="mt-4"
                block
                type="submit"
                @click="updateUser"
              >
                <v-icon left>mdi-content-save</v-icon>
                設定を更新
              </v-btn>
            </v-col>
            <v-col cols="6">
              <v-btn
                color="primary"
                class="mt-4"
                block
                @click="addNewUser"
              >
                <v-icon left>mdi-account-plus</v-icon>
                新規で追加
              </v-btn>
            </v-col>
          </v-row>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { useUserStore } from '~/store/user'

export default {
  name: 'IndexPage',
  layout: 'mobile',
  data() {
    return {
      valid: false,
      selectedUser: 1,
      users: [],
      settings: {
        id: '',
        name: '',
        gender: '',
        age: null,
        hobbies: [],
        photo: null
      },
      hobbyOptions: [
        { id: 'cstkdiat6c3011a83so0', name: '釣り' },
        { id: 'cstkdiat6c3011a83sog', name: 'キャンプ' }
      ],
      userInfo: null,
    }
  },
  mounted() {
    this.loadUserSettings()
    const userStore = useUserStore();
    userStore.initializeUser();
    this.userInfo = userStore.userInfo;
  },
  methods: {
    loadUserSettings() {
      this.fetchUserSettings(this.selectedUser).then((response) => {
        this.settings = response.data
        this.settings.hobbies = this.settings.hobbies.map(hobby =>
          this.hobbyOptions.find(option => option.id === hobby.id) || hobby
        )
      })
    },
    updateSettings() {
      if (this.$refs.form.validate()) {
        this.saveUserSettings(this.selectedUser, this.settings).then(() => {
          userStore.updateUserInfo({ id: this.selectedUser})
          alert('設定が更新されました')
        })
      }
    },
    fetchUserSettings(userId) {
      return fetch(`http://localhost:8080/users/${userId}`, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok')
          }
          return response.json()
        })
        .then(data => {
          return { data }
        })
        .catch(error => {
          console.error('There was a problem with the fetch operation:', error)
          throw error
        })
    },
    updateUser() {
      if (this.$refs.form.validate()) {
        const formData = new FormData()
        formData.append('id', this.settings.id)
        formData.append('name', this.settings.name)
        formData.append('gender', this.settings.gender)
        formData.append('age', this.settings.age)
        formData.append('hobbies', JSON.stringify(this.settings.hobbies))
        if (this.settings.photo) {
          formData.append('photo', this.settings.photo)
        }

        fetch(`http://localhost:8080/users/${this.selectedUser}`, {
          method: 'PUT',
          body: formData
        })
          .then(response => {
            if (!response.ok) {
              throw new Error('Network response was not ok')
            }
            return response.json()
          })
          .then(() => {
            alert('設定が更新されました')
          })
          .catch(error => {
            console.error('There was a problem with the fetch operation:', error)
            alert('設定の更新に失敗しました')
          })
      }
    },
    addNewUser() {
      if (this.$refs.form.validate()) {
        const formData = new FormData()
        formData.append('name', this.settings.name)
        formData.append('gender', this.settings.gender)
        formData.append('age', this.settings.age)
        formData.append('hobbies', JSON.stringify(this.settings.hobbies))
        if (this.settings.photo) {
          formData.append('photo', this.settings.photo)
        }

        fetch('http://localhost:8080/users', {
          method: 'POST',
          body: formData
        })
          .then(response => {
            if (!response.ok) {
              throw new Error('Network response was not ok')
            }
            return response.json()
          })
          .then(data => {
            alert('新規ユーザーが追加されました')
            this.users.push({ id: data.id, name: data.name })
            this.selectedUser = data.id
          })
          .catch(error => {
            console.error('There was a problem with the fetch operation:', error)
            alert('ユーザーの追加に失敗しました')
          })
      }
    }
  }
}
</script>
