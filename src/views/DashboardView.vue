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
          <!-- Ejemplo de fila -->
          <tr>
            <td>900123456</td>
            <td>Juan</td>
            <td>Pérez</td>
            <td>1234567890</td>
            <td>Nacional</td>
            <td>Natural</td>
            <td>
              Ana López - 111222333<br />
              Carlos Ruiz - 444555666
            </td>
            <td>
              Banco: Bancolombia<br />
              Cuenta: 12345678901<br />
              Tipo: Ahorros
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
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
