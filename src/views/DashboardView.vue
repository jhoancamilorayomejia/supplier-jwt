<template>
  <div class="dashboard">
    <h1>📋 Lista de Proveedores</h1>

    <div class="table-container">
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
              <div v-if="supplier.beneficiarios && supplier.beneficiarios.length">
                <div v-for="b in supplier.beneficiarios" :key="b.cedula">
                  {{ b.nombre }} - {{ b.cedula }}
                </div>
              </div>
            </td>
            <td>
              <div v-if="supplier.datos_bancarios && supplier.datos_bancarios.length">
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
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const suppliers = ref([])

onMounted(async () => {
  try {
    const res = await axios.get('/api/suppliers')
    // ahora res.data ya es un array con objetos JSON listos para iterar
    suppliers.value = res.data.map(s => ({
      ...s,
      // asegurarse de que los campos JSONB lleguen como arrays de objetos
      beneficiarios: s.beneficiarios || [],
      datos_bancarios: s.datos_bancarios || []
    }))
  } catch (error) {
    console.error('Error al obtener proveedores:', error)
  }
})
</script>

<style scoped>
.dashboard {
  min-height: 100vh;
  background: linear-gradient(135deg, #1f2937, #111827);
  padding: 30px;
  color: white;
  box-sizing: border-box;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
  font-size: 2rem;
  letter-spacing: 1px;
}

.table-container {
  width: 100%;
  height: calc(100vh - 100px);
  overflow-x: auto;
  overflow-y: auto;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 15px;
  padding: 15px;
  box-shadow: 0 0 20px rgba(0,0,0,0.5);
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
  padding: 14px 12px;
  text-align: left;
  vertical-align: top;
}

th {
  color: #38bdf8;
  font-weight: 600;
  border-bottom: 2px solid #334155;
}

tbody tr {
  transition: background 0.3s;
}

tbody tr:nth-child(even) {
  background: rgba(255,255,255,0.03);
}

tbody tr:hover {
  background: rgba(56, 189, 248, 0.15);
}

td {
  border-bottom: 1px solid #334155;
  line-height: 1.4;
  font-size: 0.95rem;
}
</style>
