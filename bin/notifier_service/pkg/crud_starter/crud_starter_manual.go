//Файл создан автоматически кодогенератором crud_generator
//Не изменяйте ничего здесь.

package crud_starter

import (
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_attachament"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_channel"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_channel_prod"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_debt_list"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_debt_list2"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_delivery_error"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_delivery_status"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_filial"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_incoming_event"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_link_type"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_mailing"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_mailing_stats"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_message"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_meter_list"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_pdf_data"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_redirect_store"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_short_links"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_statistic"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_telegram_users"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_telegram_users_info"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_template"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_template_decoration"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_u_link_store"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_w_log"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_w_log_message_del"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/crud_starter/crud_starter_w_options"

	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_attachament"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_channel"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_channel_prod"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_debt_list"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_debt_list2"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_delivery_error"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_delivery_status"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_filial"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_incoming_event"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_link_type"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_mailing"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_mailing_stats"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_message"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_meter_list"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_pdf_data"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_redirect_store"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_short_links"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_statistic"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_telegram_users"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_telegram_users_info"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_template"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_template_decoration"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_u_link_store"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_w_log"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_w_log_message_del"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/db/crud/crud_w_options"

	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_attachament"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_channel"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_channel_prod"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_debt_list"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_debt_list2"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_delivery_error"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_delivery_status"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_filial"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_incoming_event"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_link_type"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_mailing"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_mailing_stats"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_message"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_meter_list"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_pdf_data"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_redirect_store"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_short_links"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_statistic"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_telegram_users"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_telegram_users_info"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_template"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_template_decoration"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_u_link_store"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_w_log"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_w_log_message_del"
	"gitlab.aescorp.ru/dsp_dev/claim/notifier_service/pkg/network/grpc/grpc_client/grpc_w_options"
)

// initCrudTransport_manual_DB - заполняет объекты crud для работы с БД напрямую
func initCrudTransport_manual_DB() {
	crud_starter_attachament.SetCrudManualInterface(crud_attachament.Crud_DB{})
	crud_starter_channel.SetCrudManualInterface(crud_channel.Crud_DB{})
	crud_starter_channel_prod.SetCrudManualInterface(crud_channel_prod.Crud_DB{})
	crud_starter_debt_list.SetCrudManualInterface(crud_debt_list.Crud_DB{})
	crud_starter_debt_list2.SetCrudManualInterface(crud_debt_list2.Crud_DB{})
	crud_starter_delivery_error.SetCrudManualInterface(crud_delivery_error.Crud_DB{})
	crud_starter_delivery_status.SetCrudManualInterface(crud_delivery_status.Crud_DB{})
	crud_starter_filial.SetCrudManualInterface(crud_filial.Crud_DB{})
	crud_starter_incoming_event.SetCrudManualInterface(crud_incoming_event.Crud_DB{})
	crud_starter_link_type.SetCrudManualInterface(crud_link_type.Crud_DB{})
	crud_starter_mailing.SetCrudManualInterface(crud_mailing.Crud_DB{})
	crud_starter_mailing_stats.SetCrudManualInterface(crud_mailing_stats.Crud_DB{})
	crud_starter_message.SetCrudManualInterface(crud_message.Crud_DB{})
	crud_starter_meter_list.SetCrudManualInterface(crud_meter_list.Crud_DB{})
	crud_starter_pdf_data.SetCrudManualInterface(crud_pdf_data.Crud_DB{})
	crud_starter_redirect_store.SetCrudManualInterface(crud_redirect_store.Crud_DB{})
	crud_starter_short_links.SetCrudManualInterface(crud_short_links.Crud_DB{})
	crud_starter_statistic.SetCrudManualInterface(crud_statistic.Crud_DB{})
	crud_starter_telegram_users.SetCrudManualInterface(crud_telegram_users.Crud_DB{})
	crud_starter_telegram_users_info.SetCrudManualInterface(crud_telegram_users_info.Crud_DB{})
	crud_starter_template.SetCrudManualInterface(crud_template.Crud_DB{})
	crud_starter_template_decoration.SetCrudManualInterface(crud_template_decoration.Crud_DB{})
	crud_starter_u_link_store.SetCrudManualInterface(crud_u_link_store.Crud_DB{})
	crud_starter_w_log.SetCrudManualInterface(crud_w_log.Crud_DB{})
	crud_starter_w_log_message_del.SetCrudManualInterface(crud_w_log_message_del.Crud_DB{})
	crud_starter_w_options.SetCrudManualInterface(crud_w_options.Crud_DB{})
}

// initCrudTransport_manual_GRPC - заполняет объекты crud для работы с БД через протокол GRPC
func initCrudTransport_manual_GRPC() {
	crud_starter_attachament.SetCrudManualInterface(grpc_attachament.Crud_GRPC{})
	crud_starter_channel.SetCrudManualInterface(grpc_channel.Crud_GRPC{})
	crud_starter_channel_prod.SetCrudManualInterface(grpc_channel_prod.Crud_GRPC{})
	crud_starter_debt_list.SetCrudManualInterface(grpc_debt_list.Crud_GRPC{})
	crud_starter_debt_list2.SetCrudManualInterface(grpc_debt_list2.Crud_GRPC{})
	crud_starter_delivery_error.SetCrudManualInterface(grpc_delivery_error.Crud_GRPC{})
	crud_starter_delivery_status.SetCrudManualInterface(grpc_delivery_status.Crud_GRPC{})
	crud_starter_filial.SetCrudManualInterface(grpc_filial.Crud_GRPC{})
	crud_starter_incoming_event.SetCrudManualInterface(grpc_incoming_event.Crud_GRPC{})
	crud_starter_link_type.SetCrudManualInterface(grpc_link_type.Crud_GRPC{})
	crud_starter_mailing.SetCrudManualInterface(grpc_mailing.Crud_GRPC{})
	crud_starter_mailing_stats.SetCrudManualInterface(grpc_mailing_stats.Crud_GRPC{})
	crud_starter_message.SetCrudManualInterface(grpc_message.Crud_GRPC{})
	crud_starter_meter_list.SetCrudManualInterface(grpc_meter_list.Crud_GRPC{})
	crud_starter_pdf_data.SetCrudManualInterface(grpc_pdf_data.Crud_GRPC{})
	crud_starter_redirect_store.SetCrudManualInterface(grpc_redirect_store.Crud_GRPC{})
	crud_starter_short_links.SetCrudManualInterface(grpc_short_links.Crud_GRPC{})
	crud_starter_statistic.SetCrudManualInterface(grpc_statistic.Crud_GRPC{})
	crud_starter_telegram_users.SetCrudManualInterface(grpc_telegram_users.Crud_GRPC{})
	crud_starter_telegram_users_info.SetCrudManualInterface(grpc_telegram_users_info.Crud_GRPC{})
	crud_starter_template.SetCrudManualInterface(grpc_template.Crud_GRPC{})
	crud_starter_template_decoration.SetCrudManualInterface(grpc_template_decoration.Crud_GRPC{})
	crud_starter_u_link_store.SetCrudManualInterface(grpc_u_link_store.Crud_GRPC{})
	crud_starter_w_log.SetCrudManualInterface(grpc_w_log.Crud_GRPC{})
	crud_starter_w_log_message_del.SetCrudManualInterface(grpc_w_log_message_del.Crud_GRPC{})
	crud_starter_w_options.SetCrudManualInterface(grpc_w_options.Crud_GRPC{})
}

// initCrudTransport_manual_NRPC - заполняет объекты crud для работы с БД через протокол NRPC
func initCrudTransport_manual_NRPC() {
	crud_starter_attachament.SetCrudManualInterface(grpc_attachament.Crud_GRPC{})
	crud_starter_channel.SetCrudManualInterface(grpc_channel.Crud_GRPC{})
	crud_starter_channel_prod.SetCrudManualInterface(grpc_channel_prod.Crud_GRPC{})
	crud_starter_debt_list.SetCrudManualInterface(grpc_debt_list.Crud_GRPC{})
	crud_starter_debt_list2.SetCrudManualInterface(grpc_debt_list2.Crud_GRPC{})
	crud_starter_delivery_error.SetCrudManualInterface(grpc_delivery_error.Crud_GRPC{})
	crud_starter_delivery_status.SetCrudManualInterface(grpc_delivery_status.Crud_GRPC{})
	crud_starter_filial.SetCrudManualInterface(grpc_filial.Crud_GRPC{})
	crud_starter_incoming_event.SetCrudManualInterface(grpc_incoming_event.Crud_GRPC{})
	crud_starter_link_type.SetCrudManualInterface(grpc_link_type.Crud_GRPC{})
	crud_starter_mailing.SetCrudManualInterface(grpc_mailing.Crud_GRPC{})
	crud_starter_mailing_stats.SetCrudManualInterface(grpc_mailing_stats.Crud_GRPC{})
	crud_starter_message.SetCrudManualInterface(grpc_message.Crud_GRPC{})
	crud_starter_meter_list.SetCrudManualInterface(grpc_meter_list.Crud_GRPC{})
	crud_starter_pdf_data.SetCrudManualInterface(grpc_pdf_data.Crud_GRPC{})
	crud_starter_redirect_store.SetCrudManualInterface(grpc_redirect_store.Crud_GRPC{})
	crud_starter_short_links.SetCrudManualInterface(grpc_short_links.Crud_GRPC{})
	crud_starter_statistic.SetCrudManualInterface(grpc_statistic.Crud_GRPC{})
	crud_starter_telegram_users.SetCrudManualInterface(grpc_telegram_users.Crud_GRPC{})
	crud_starter_telegram_users_info.SetCrudManualInterface(grpc_telegram_users_info.Crud_GRPC{})
	crud_starter_template.SetCrudManualInterface(grpc_template.Crud_GRPC{})
	crud_starter_template_decoration.SetCrudManualInterface(grpc_template_decoration.Crud_GRPC{})
	crud_starter_u_link_store.SetCrudManualInterface(grpc_u_link_store.Crud_GRPC{})
	crud_starter_w_log.SetCrudManualInterface(grpc_w_log.Crud_GRPC{})
	crud_starter_w_log_message_del.SetCrudManualInterface(grpc_w_log_message_del.Crud_GRPC{})
	crud_starter_w_options.SetCrudManualInterface(grpc_w_options.Crud_GRPC{})
}