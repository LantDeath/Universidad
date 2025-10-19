install.packages("readxl")  # Instalar el paquete para leer archivos .xls
library(readxl)
install.packages("ggplot2") # Paquete para graficar.
library(ggplot2)


# Instalamos y creamos dos dataframe con los datasets obtenidos sobre el consumo y producción energético.

Elec_generation <- read_excel("electricity-generation.xlsx")
Elec_consumption <- read_excel("energy_consumption.xls")



# ¿Es el consumo energético promedio diferente de la producción energética promedio? 
# Calculamos las medias.
mean_generation <- mean(Elec_generation$`Electricity generation - TWh`)
mean_consumption <- mean(Elec_consumption$`Total Energy Consumption (TWh)`)

# Usamos una prueba para comparar las medias.
result_test <- t.test(Elec_generation$`Electricity generation - TWh`, Elec_consumption$`Total Energy Consumption (TWh)` )

# Mostramos los resultados.
print(paste("Media de consumo:", mean_generation))
print(paste("Media de producción:", mean_consumption))
print(result_test)




#¿Cuál es la dispersión de los datos en el consumo y la producción energética?

# Generamos variables y la relación entre las variables. (Varianza y desvación estándar)
var_consumo <- var(Elec_generation$`Electricity generation - TWh`)
sd_consumo <- sd(Elec_generation$`Electricity generation - TWh`)

var_produccion <- var(Elec_consumption$`Total Energy Consumption (TWh)`)
sd_produccion <- sd(Elec_consumption$`Total Energy Consumption (TWh)`)

# Mostrar los resultados.
print(paste("Varianza de consumo:", var_consumo))
print(paste("Desviación estándar de consumo:", sd_consumo))
print(paste("Varianza de producción:", var_produccion))
print(paste("Desviación estándar de producción:", sd_produccion))



# CREAMOS LAS GRÁFICAS DE NUESTROS DATOS.



# Crear un dataframe con las medias
medias <- data.frame(
  Tipo = c("Generación", "Consumo"),
  Media = c(mean_generation, mean_consumption)
)

# Graficas de las medias
ggplot(medias, aes(x = Tipo, y = Media, fill = Tipo)) +
  geom_bar(stat = "identity", position = "dodge") +
  labs(title = "Comparación de Medias de Generación y Consumo Energético",
       x = "Tipo",
       y = "Media (TWh)") +
  theme_minimal()

# Histograma para la generación de energía
hist(Elec_generation$`Electricity generation - TWh`, 
     main = "Distribución de Generación Energética",
     xlab = "Generación (TWh)",
     ylab = "Frecuencia",
     col = "lightblue",
     border = "black")

# Histograma para el consumo de energía
hist(Elec_consumption$`Total Energy Consumption (TWh)`, 
     main = "Distribución de Consumo Energético",
     xlab = "Consumo (TWh)",
     ylab = "Frecuencia",
     col = "lightgreen",
     border = "black")