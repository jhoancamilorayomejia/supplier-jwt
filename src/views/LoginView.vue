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
  width: 480px;              /* 🔥 MÁS GRANDE */
  padding: 60px 50px;       /* 🔥 MÁS ESPACIADO */
  border-radius: 25px;
  background: rgba(255, 255, 255, 0.08);
  backdrop-filter: blur(20px);
  box-shadow: 0 0 60px rgba(0, 0, 0, 0.8);
  color: white;
  text-align: center;
}

.login-box h1 {
  font-size: 2.8rem;        /* 🔥 TÍTULO MÁS GRANDE */
  margin-bottom: 10px;
  letter-spacing: 3px;
}

.login-box p {
  color: #bbb;
  margin-bottom: 40px;
  font-size: 1.1rem;
}

.field {
  position: relative;
  margin-bottom: 35px;      /* MÁS SEPARACIÓN */
}

.field input {
  width: 100%;
  padding: 18px 12px;       /* INPUT MÁS GRANDE */
  background: transparent;
  border: none;
  border-bottom: 2px solid #aaa;
  color: white;
  font-size: 1.2rem;
}

.field label {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  transition: 0.3s;
  color: #aaa;
  pointer-events: none;
  font-size: 1.1rem;
}

.field input:focus + label,
.field input:not(:placeholder-shown) + label {
  top: -10px;
  font-size: 0.85rem;
  color: #4fc3f7;
}

button {
  width: 100%;
  padding: 18px;
  margin-top: 20px;
  border-radius: 40px;
  background: linear-gradient(45deg, #4fc3f7, #00e5ff);
  border: none;
  color: #000;
  font-weight: bold;
  font-size: 1.1rem;
  cursor: pointer;
  letter-spacing: 2px;
  transition: 0.3s;
}

button:hover {
  transform: scale(1.08);
}

.mensaje {
  margin-top: 20px;
  font-size: 1rem;
}

</style>
