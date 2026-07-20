{{/*
Возвращает имя приложения.
Например:
dev-app
prod-app
*/}}
{{- define "efftest.appName" -}}
{{- printf "%s-app" .Release.Name -}}
{{- end -}}

{{/*
Возвращает имя бд.
Например:
dev-postgres
prod-postgres
*/}}
{{- define "efftest.postgresName" -}}
{{ printf "%s-postgres" .Release.Name }}
{{- end }}

{{/*
Labels, используемые в selector.
Их нельзя менять после создания Deployment.
*/}}
{{- define "efftest.selectorLabels" -}}
app: {{ include "efftest.appName" . }}
{{- end }}

{{- define "efftest.postgresSelectorLabels" -}}
app: {{ include "efftest.postgresName" . }}
{{- end }}

{{/*
Полный набор labels для всех ресурсов.
*/}}
{{- define "efftest.labels" -}}
{{ include "efftest.selectorLabels" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "efftest.configMapName" -}}
{{ printf "%s-config" .Release.Name }}
{{- end }}

{{- define "efftest.secretName" -}}
{{ printf "%s-secret" .Release.Name }}
{{- end }}

