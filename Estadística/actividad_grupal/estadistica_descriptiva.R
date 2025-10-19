#Librerías necesarias
library(readxl)
library(dplyr)

datos <- read_excel(file.choose("Compraventa de viviendas.xls"))

#Limpia y prepara los datos
datos_limpios <- datos %>%
  rename(Año = 1, Mes = 2, Datos = 3) %>%  #Renombra las columnas
  filter(!is.na(Datos)) %>%                #Elimina las filas con valores nulos
  mutate(
    Datos = as.numeric(sub(",", ".", Datos)),  
    Año = as.numeric(Año)
  ) %>%
  filter(Año >= 2020 & Año <= 2024)           #Filtra por el rango de años

#Calcula las estadísticas descriptivas
estadisticas <- datos_limpios %>%
  summarise(
    Media = mean(Datos),
    Varianza = var(Datos),
    Desviacion_Tipica = sd(Datos),
    Minimo = min(Datos),
    Maximo = max(Datos),
    Observaciones = n()
  )

#Muestra los resultados
cat("Media, varianza y desviación típica (2020-2024)")
print(estadisticas)

#Calcula la media porcada año de manera individual
estadisticas_por_año <- datos_limpios %>%
  group_by(Año) %>%
  summarise(
    Media = mean(Datos))

cat("Media por años (2020-2024)")
print(estadisticas_por_año)

#Gráfico de medias anuales
ggplot(estadisticas_por_año, aes(x = factor(Año), y = Media)) +
  geom_col(fill = "lightgreen", color = "black", width = 0.6) + 
  geom_text(aes(label = round(Media, 1)), vjust = -0.5, size = 4) + 
  labs(title = "Media Anual de Compraventa de Viviendas",
       subtitle = "Periodo 2020-2024",
       x = "Años",
       y = "Número medio de viviendas (en miles)") +
  theme_minimal() +
  theme(plot.title = element_text(hjust = 0.5, face = "bold"),  
        plot.subtitle = element_text(hjust = 0.5))