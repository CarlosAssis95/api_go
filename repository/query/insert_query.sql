INSERT INTO Dados (
    ds_operacao, ds_paciente, cd_paciente, dt_nascimento, ds_sexo, ds_accession_number, 
    nr_identificador, ds_procedimento, ds_medico, ds_crm_nr, ds_crm_uf, bb_laudo, 
    bb_rtf, dt_assinatura, dt_data, medico_solicitante, codigo_procedimento, tipo_exame, modalidade
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
);