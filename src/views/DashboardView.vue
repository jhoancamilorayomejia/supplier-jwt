<template>
  <div class="dashboard">
    <h1>📋 Lista de Proveedores</h1>

    <!-- Mostrar token -->
    <!--div v-if="token" class="token">
      <strong>Token:</strong> {{ token }}
    </div-->

    <!-- Botón de cerrar sesión -->
    <!-- Botón de cerrar sesión -->
<div class="logout">
  <button @click="logout">Salir</button>
</div>


    <!-- Mostrar error -->
    <div v-if="error" class="error">{{ error }}</div>

    <!-- Tabla de proveedores -->
    <div class="table-container" v-if="!loading && suppliers.length">
      <table>
        <thead>
          <tr>
            <th>NIT</th>
            <th>Nombre</th>
            <th>Apellido</th>
            <th>Cédula</th>
            <th>Tipo Proveedor</th>
            <th>Tipo Persona</th>
            <th>Beneficiarios / Socios</th>
            <th>Datos Bancarios</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="supplier in suppliers" :key="supplier.id">
            <td>{{ supplier.nit }}</td>
            <td>{{ supplier.nombre }}</td>
            <td>{{ supplier.apellido }}</td>
            <td>{{ supplier.cedula }}</td>
            <td>{{ supplier.tipo_proveedor }}</td>
            <td>{{ supplier.tipo_persona }}</td>

            <td>
              <div v-if="supplier.beneficiarios.length">
                <div v-for="b in supplier.beneficiarios" :key="b.cedula">
                  {{ b.nombre }} - {{ b.cedula }}
                </div>
              </div>
            </td>

            <td>
              <div v-if="supplier.datos_bancarios.length">
                <div v-for="(d, index) in supplier.datos_bancarios" :key="d.num_cuenta">
                  Banco: {{ d.banco }}<br />
                  Cuenta: {{ d.num_cuenta }}<br />
                  Tipo: {{ d.tipo_cuenta }}
                  <hr v-if="index !== supplier.datos_bancarios.length - 1" />
                </div>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <p v-if="loading">Cargando proveedores...</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const suppliers = ref([])
const loading = ref(true)
const error = ref('')
const token = ref('')
const router = useRouter()

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/api/login') // redirige al login
}

onMounted(async () => {
  const savedToken = localStorage.getItem('token')
  token.value = savedToken

  if (!savedToken) {
    router.push('/api/login')
    return
  }

  try {
    const res = await axios.get('http://localhost:8080/api/suppliers', {
      headers: {
        Authorization: `Bearer ${savedToken}`
      }
    })

    suppliers.value = res.data.map(s => ({
      ...s,
      beneficiarios: typeof s.beneficiarios === 'string' ? JSON.parse(s.beneficiarios) : (s.beneficiarios || []),
      datos_bancarios: typeof s.datos_bancarios === 'string' ? JSON.parse(s.datos_bancarios) : (s.datos_bancarios || [])
    }))

  } catch (err) {
    if (err.response?.status === 401) {
      logout()
    } else {
      error.value = 'Error al cargar proveedores'
      console.error(err)
    }
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.dashboard {
  min-height: 100vh;
  background: linear-gradient(135deg, #1f2937, #111827);
  padding: 30px;
  color: white;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

.token {
  background: #0ea5e9;
  padding: 10px;
  border-radius: 8px;
  margin-bottom: 10px;
  word-break: break-all;
  text-align: center;
}

.logout {
  position: absolute;   /* Para posicionarlo encima del contenido */
  top: 20px;            /* Separación desde arriba */
  right: 20px;          /* Separación desde la derecha */
}

.logout button {
  padding: 10px 20px;
  border: none;
  border-radius: 20px;
  background: black;    /* Cambiado a negro */
  color: white;
  font-weight: bold;
  cursor: pointer;
  transition: 0.3s;
}

.logout button:hover {
  background: #333;     /* Oscurece un poco al pasar el mouse */
}


.error {
  background: #dc2626;
  padding: 10px;
  border-radius: 8px;
  margin-bottom: 15px;
  text-align: center;
}

.table-container {
  overflow-x: auto;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 15px;
  padding: 15px;
}

table {
  width: 100%;
  border-collapse: collapse;
  min-width: 1200px;
}

thead {
  background: #0f172a;
}

th, td {
  padding: 12px;
  border-bottom: 1px solid #334155;
}

th {
  color: #38bdf8;
}
</style>
