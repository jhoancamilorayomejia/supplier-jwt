<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()

const username = ref('')
const password = ref('')
const mensaje = ref('')
const loading = ref(false)

const login = async () => {
  mensaje.value = ''
  loading.value = true

  try {
    const response = await axios.post('http://localhost:8080/api/login', {
      username: username.value,
      password: password.value
    })

    mensaje.value = response.data.message || '✅ Login exitoso'
    router.push('/dashboard')

  } catch  {
    mensaje.value = '❌ Usuario o contraseña incorrectos'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="screen">
    <div class="login-box">
      <h1>Login</h1>
      <p>Acceso al sistema</p>

      <form @submit.prevent="login">
        <div class="field">
          <input
            type="text"
            v-model="username"
            required
            placeholder=" "
          />
          <label>Usuario</label>
        </div>

        <div class="field">
          <input
            type="password"
            v-model="password"
            required
            placeholder=" "
          />
          <label>Contraseña</label>
        </div>

        <button :disabled="loading">
          {{ loading ? 'Ingresando...' : 'INGRESAR' }}
        </button>

        <p v-if="mensaje" class="mensaje">{{ mensaje }}</p>
      </form>
    </div>
  </div>
</template>

<style scoped>
.screen {
  min-height: 100vh;
  background: radial-gradient(circle at top, #0f2027, #203a43, #2c5364);
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: 'Segoe UI', sans-serif;
}

.login-box {
  width: 380px;
  padding: 40px;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.08);
  backdrop-filter: blur(15px);
  box-shadow: 0 0 40px rgba(0, 0, 0, 0.6);
  color: white;
  text-align: center;
}

.login-box h1 {
  font-size: 2rem;
  margin-bottom: 5px;
  letter-spacing: 2px;
}

.login-box p {
  color: #aaa;
  margin-bottom: 30px;
}

.field {
  position: relative;
  margin-bottom: 25px;
}

.field input {
  width: 100%;
  padding: 14px 10px;
  background: transparent;
  border: none;
  border-bottom: 2px solid #aaa;
  color: white;
  font-size: 1rem;
}

.field label {
  position: absolute;
  left: 10px;
  top: 50%;
  transform: translateY(-50%);
  transition: 0.3s;
  color: #aaa;
  pointer-events: none;
}

.field input:focus + label,
.field input:not(:placeholder-shown) + label {
  top: -8px;
  font-size: 0.8rem;
  color: #4fc3f7;
}

.field input:focus {
  outline: none;
  border-bottom: 2px solid #4fc3f7;
}

button {
  width: 100%;
  padding: 14px;
  margin-top: 15px;
  border-radius: 30px;
  background: linear-gradient(45deg, #4fc3f7, #00e5ff);
  border: none;
  color: #000;
  font-weight: bold;
  cursor: pointer;
  letter-spacing: 1px;
  transition: 0.3s;
}

button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

button:hover {
  transform: scale(1.05);
}

.mensaje {
  margin-top: 15px;
  font-size: 0.9rem;
}
</style>
