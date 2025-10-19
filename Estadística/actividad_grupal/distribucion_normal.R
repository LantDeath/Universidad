#Librerías necesarias
library(ggplot2)
library(readxl)
library(dplyr)
library(patchwork)

#Función para cargar los datos
cargar_datos <- function() {
  if(file.exists("Compraventa de viviendas.xls")) {
    return(read_excel("Compraventa de viviendas.xls"))
  }
  message("Archivo 'Compraventa de viviendas.xls' no encontrado en el directorio actual.")
  message("Por favor selecciona el archivo manualmente...")
  return(read_excel(file.choose()))
}

#Lee los datos históricos
datos <- cargar_datos()

#Limpia y prepara los datos históricos
datos_limpios <- datos %>%
  rename(Año = 1, Mes = 2, Datos = 3) %>%
  filter(!is.na(Datos)) %>%
  mutate(Datos = as.numeric(sub(",", ".", Datos)),
         Año = as.factor(Año))

#Añadimos manualmente los datos de 2025 (enero y febrero)
datos_2025 <- data.frame(
  Año = as.factor(2025),
  Mes = c("Enero", "Febrero"),
  Datos = c(60.651, 59.682)
)

#Lo combinamos con datos históricos
datos_completos <- bind_rows(datos_limpios, datos_2025) %>%
  mutate(Año = as.factor(Año))

#Gráficos separados por año (incluyendo el 2025)
por_año_plot <- ggplot(datos_completos, aes(x = Datos)) +
  geom_histogram(aes(y = ..density..), 
                 bins = 8,
                 fill = "lightgreen", 
                 color = "black") +
  stat_function(fun = dnorm, 
                args = list(mean = mean(datos_completos$Datos), 
                            sd = sd(datos_completos$Datos)),
                color = "blue", 
                size = 0.8,
                linetype = "dashed") +
  facet_wrap(~Año, scales = "free_y", ncol = 3) +
  labs(subtitle = "Distribución por Año (Incluyendo datos parciales 2025)",
       x = "Número de viviendas (en miles)",
       y = "Densidad") +
  theme_minimal()

# Combinamos ambos gráficos
combinado_plot <- por_año_plot + 
  plot_annotation(title = "Análisis de Distribución Normal de Compraventa de Viviendas")

#Muestra el gráfico combinado
print(combinado_plot)

#Calculamos las estadísticas por año (incluyendo 2025)
estadisticas_año <- datos_completos %>%
  group_by(Año) %>%
  summarise(Media = mean(Datos),
            Desviación = sd(Datos),
            Mínimo = min(Datos),
            Máximo = max(Datos),
            Observaciones = n(),
            .groups = 'drop')

estadisticas_año <- estadisticas_año %>%
  mutate(Nota = ifelse(Año == 2025, 
                       "Estimación basada en datos parciales (Enero-Febrero)", 
                       "Datos completos"))

print(estadisticas_año)